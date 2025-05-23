package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Step 2 in setting up the backend

func SetupDatabase(user string, password string, dbname string, host string, port string) (*sql.DB, error){
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