package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SomeKay/recipe-api/db"
	"github.com/SomeKay/recipe-api/handlers"
	_ "github.com/lib/pq"
)

func main() {
	database, err := db.ConnectToDb()
	if err != nil {
		log.Fatalf("ConnectToDb error: %v", err)
	}
	fmt.Println("Successfully connected to db!")

	db.CreateUsersTable(database)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	userHandler := &handlers.UserHandler{Db: database}

	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/users/", userHandler.GetUser)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on port 8080")

	defer database.Close()
}
