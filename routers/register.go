package routers

import (
	"encoding/json"
	"net/http"

	"github.com/nelsomar/twittor/bd"
	"github.com/nelsomar/twittor/models"
)

/* Registro, function to crate users register */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Users
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error on data recived: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password length must be greater than 6", 400)
		return
	}

	_, foundUser, _ := bd.CheckUserExists(t.Email)

	if !foundUser {
		http.Error(w, "User exists, please try again", 400)
		return
	}

	_, status, err := bd.RegisterUser(t)

	if err != nil {
		http.Error(w, "Error registering user: "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Error registering user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
