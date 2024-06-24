package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/models"
	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/storage"
)

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Message string
		Error   string
	}{
		Title: "Withdraw",
	}

	if r.Method == "POST" {
		amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
		if err != nil || amount <= 0 {
			data.Error = "Please enter a valid positive amount"
		} else {
			session, _ := store.Get(r, "session")
			username, ok := session.Values["username"].(string)
			if !ok {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			db := storage.GetDB()
			account, err := models.GetAccountByUsername(db, username)
			if err != nil {
				data.Error = "Failed to get account"
			} else {
				err = account.Withdraw(db, amount)
				if err != nil {
					data.Error = "Failed to withdraw: " + err.Error()
				} else {
					data.Message = "Withdrawal successful"
				}
			}
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/withdraw.html"))
	tmpl.Execute(w, data)
}