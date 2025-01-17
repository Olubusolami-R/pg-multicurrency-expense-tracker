package repository

import (
	"database/sql"
	"fmt"
	"strings"

	models "github.com/Olubusolami-R/multicurrency-tracker/internal/models"
	"github.com/lib/pq"
)

type CurrencyRepository interface {
	CreateSingleCurrency(currency models.Currency) error
	CreateMultipleCurrencies(currencies []models.Currency) error
	GetCurrencies()([]models.Currency, error)
	GetCurrencyIDsByCode(codes []string)(map[string]uint, error)
	CheckCurrenciesPopulated() (bool,error)
}

// Handles database operations for currency
type currencyRepository struct {
	DB *sql.DB
}

// Helper function to initialize the repository
func NewCurrencyRepository(db *sql.DB) CurrencyRepository{
	return &currencyRepository{DB:db}
}

func (r *currencyRepository) CheckCurrenciesPopulated() (bool,error){
	var count int
	query := "SELECT COUNT(*) from currencies"
	err := r.DB.QueryRow(query).Scan(&count)
    if err != nil {
        return false, fmt.Errorf("error checking currencies table: %w", err)
    }
    return count > 0, nil
}

func (r *currencyRepository) CreateSingleCurrency(currency models.Currency) error {

	query := "INSERT INTO currencies (code, name) VALUES ($1, $2)"
	_,err := r.DB.Exec(query, currency.Code, currency.Name)
	if err != nil {
		return err
	}
	return nil

}

// Batch inserts
func (r *currencyRepository) CreateMultipleCurrencies(currencies []models.Currency) error {

	query := "INSERT INTO currencies (code, name) VALUES "
	values := []interface{}{}
	placeholders := []string{}

	for i, currency := range currencies{
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		values = append(values, currency.Code, currency.Name)
	}

	query += strings.Join(placeholders, ",")


	_, err := r.DB.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (r *currencyRepository) GetCurrencies()([]models.Currency, error){

	query:="SELECT code, name FROM currencies"

	
	rows,err:=r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch currencies: %w", err)
	}
	defer rows.Close()

	var currencies []models.Currency

	for rows.Next() {

		var currency models.Currency

		if err := rows.Scan(&currency.Code, &currency.Name); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		currencies = append(currencies, currency)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after row iteration: %w", err)
	}

	return currencies, nil
}

func (r *currencyRepository) GetCurrencyIDsByCode(codes []string)(map[string]uint, error){
	
	query:= "SELECT id, code, name FROM currencies WHERE code=ANY($1)"

	rows,err:=r.DB.Query(query,pq.Array(&codes))
	if err != nil {
		return nil, fmt.Errorf("error querying currencies: %w", err)
	}
	defer rows.Close()

	// Parse the rows of the matched currencies and get it into a list of currencies (objects)
	var currencies []models.Currency
	for rows.Next(){
		var currency models.Currency
		if err := rows.Scan(&currency.ID, &currency.Code, &currency.Name); err!=nil{
			return nil,fmt.Errorf("failed to scan currency row")
		}
		currencies=append(currencies, currency)
	}

	if err:=rows.Err();err!=nil{
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	//Make the list of currencies accessible by converting to map
	currencyMap := make(map[string]uint)
	for _,currency:=range(currencies){
		currencyMap[currency.Code]=currency.ID
	}

	return currencyMap, nil
}