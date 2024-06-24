package main

import (
	"log"
	"net/http"
	"os"

	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/handlers"
	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/storage"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	store *sessions.CookieStore
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	storage.InitDB()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/signup", handlers.SignupHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/deposit", handlers.DepositHandler).Methods("GET", "POST")
	r.HandleFunc("/withdraw", handlers.WithdrawHandler).Methods("GET", "POST")
	r.HandleFunc("/upload", handlers.UploadHandler).Methods("GET", "POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", r)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}