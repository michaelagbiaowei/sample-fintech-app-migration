package handlers

import (
	"html/template"
	"net/http"

	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/auth"
	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/storage"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Message string
		Error   string
	}{
		Title: "Login",
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			data.Error = "Username and password are required"
		} else {
			db := storage.GetDB()
			authenticated, err := auth.AuthenticateUser(db, username, password)
			if err != nil {
				data.Error = "Authentication failed: " + err.Error()
			} else if authenticated {
				session, _ := store.Get(r, "session")
				session.Values["authenticated"] = true
				session.Values["username"] = username
				session.Save(r, w)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			} else {
				data.Error = "Invalid username or password"
			}
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/login.html"))
	tmpl.Execute(w, data)
}