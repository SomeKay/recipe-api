package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SomeKay/recipe-api/models"
)

func CreateUsersTable(db *sql.DB) {
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);
	`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to create users table: %v", err)
	}

	fmt.Println("Successfully created users table")
}

func InsertUser(db *sql.DB, user models.User) error {
	sqlStatement := `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
	`

	_, err := db.Exec(sqlStatement, user.Name, user.Email)
	if err != nil {
		return fmt.Errorf("unable to insert user: %v", err)
	}
	fmt.Println("Successfully inserted user")
	return nil
}

func GetUser(db *sql.DB, id int) (models.User, error) {
	sqlStatement := `
		SELECT * FROM users WHERE id=$1;
	`

	var user models.User
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("no user found with id %d", id)
		}
		return user, fmt.Errorf("error with query: %v", err)
	}

	return user, nil
}
