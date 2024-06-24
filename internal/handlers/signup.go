package handlers

import (
	"html/template"
	"net/http"

	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/auth"
	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/models"
	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/storage"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Message string
		Error   string
	}{
		Title: "Sign Up",
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			data.Error = "Username and password are required"
		} else {
			db := storage.GetDB()
			err := auth.RegisterUser(db, username, password)
			if err != nil {
				data.Error = "Failed to register user: " + err.Error()
			} else {
				// Create an account for the new user
				user, err := models.GetUserByUsername(db, username)
				if err != nil {
					data.Error = "Failed to get user: " + err.Error()
				} else {
					err = models.CreateAccount(db, user.ID)
					if err != nil {
						data.Error = "Failed to create account: " + err.Error()
					} else {
						data.Message = "User registered successfully. Please log in."
					}
				}
			}
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/signup.html"))
	tmpl.Execute(w, data)
}