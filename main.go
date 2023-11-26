package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SomeKay/recipe-api/db"
	"github.com/SomeKay/recipe-api/models"
	_ "github.com/lib/pq"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5432, "postgres", "postgres", "test")

	fmt.Println("Database string: ", psqlInfo)
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("sql Open error: ", err)
	}
	defer database.Close()

	err = database.Ping()
	if err != nil {
		log.Fatal("Ping error: ", err)
	}

	fmt.Println("Successfully connected to db!")

	db.CreateUsersTable(database)

	user := models.User{Name: "Bruce Wayne", Email: "bruce.wayne@gotham.com"}
	if err := db.InsertUser(database, user); err != nil {
		log.Printf("insertUser error: %v", err)
	}

	user, err = db.GetUser(database, 1)
	if err != nil {
		log.Printf("getUser error: %v", err)
	} else {
		fmt.Println("User: ", user)
	}
}
