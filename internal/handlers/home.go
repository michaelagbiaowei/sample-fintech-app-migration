package handlers

import (
	"html/template"
	"net/http"

	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/models"
	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/storage"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username, ok := session.Values["username"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	db := storage.GetDB()
	account, err := models.GetAccountByUsername(db, username)
	if err != nil {
		http.Error(w, "Failed to get account", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title    string
		Username string
		Balance  float64
	}{
		Title:    "Home",
		Username: username,
		Balance:  account.Balance,
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/home.html"))
	tmpl.Execute(w, data)
}