package handlers

import (
    "html/template"
    "net/http"
    "strconv"

    "github.com/gorilla/sessions"
    "github.com/michaelagbiaowei/sample-fintech-app-migration/internal/models"
    "github.com/michaelagbiaowei/sample-fintech-app-migration/internal/storage"
)

var store *sessions.CookieStore

func init() {
    // Replace "your-secret-key" with a actual secret key
    store = sessions.NewCookieStore([]byte("your-secret-key"))
}

func DepositHandler(w http.ResponseWriter, r *http.Request) {
    data := struct {
        Title   string
        Message string
        Error   string
    }{
        Title: "Deposit",
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
                err = account.Deposit(db, amount)
                if err != nil {
                    data.Error = "Failed to deposit"
                } else {
                    data.Message = "Deposit successful"
                }
            }
        }
    }

    tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/deposit.html"))
    tmpl.Execute(w, data)
}