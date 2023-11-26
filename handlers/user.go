package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SomeKay/recipe-api/db"
	"github.com/SomeKay/recipe-api/models"
)

type UserHandler struct {
	Db *sql.DB
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)
		if err := db.InsertUser(h.Db, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, _ := strconv.Atoi(r.URL.Path[len("/users/"):])
		user, err := db.GetUser(h.Db, id)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}
