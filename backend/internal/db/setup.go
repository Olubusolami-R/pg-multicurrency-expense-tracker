package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func SetupDatabase(user, password, dbname, host string, port int) (*sql.DB, error){
	connStr:=fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", 
	user, password, dbname, host, port)

	db,err:=sql.Open("postgres",connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database is unreachable: %w", err)
	}

	log.Println("Database connected successfully.")
	return db, nil
}