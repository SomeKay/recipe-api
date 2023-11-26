package db

import (
	"database/sql"
	"fmt"
)

func ConnectToDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5432, "postgres", "postgres", "test")

	fmt.Println("Database string: ", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("sql Open error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping error: %v", err)
	}

	return db, nil
}
