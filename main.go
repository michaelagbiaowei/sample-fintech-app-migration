package main

import (
	"log"
	"net/http"
	"os"

	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/handlers"
	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/storage"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore
)

func init() {
	// Use environment variable directly, with a default value if not set
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		sessionKey = "default_session_key" // It's better to generate this randomly in a production environment
	}

	store = sessions.NewCookieStore([]byte(sessionKey))
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
}
