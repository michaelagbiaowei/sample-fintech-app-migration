package tests

import (
	"database/sql"
	"testing"

	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/auth"
	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/storage"
)

func TestRegisterAndAuthenticateUser(t *testing.T) {
	t.Log("Initializing database...")
	storage.InitDB()
	db := storage.GetDB()
	if db == nil {
		t.Fatal("Failed to get database connection")
	}

	username := "testuser"
	password := "testpassword"

	// Clean up any existing test user before starting
	cleanupUser(t, db, username)

	t.Log("Registering user...")
	err := auth.RegisterUser(db, username, password)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	t.Log("Checking if user was inserted...")
	if !userExists(t, db, username) {
		t.Fatal("User was not inserted into the database")
	}

	t.Log("Authenticating user...")
	authenticated, err := auth.AuthenticateUser(db, username, password)
	if err != nil {
		t.Fatalf("Error during authentication: %v", err)
	}

	if !authenticated {
		t.Error("Expected user to be authenticated, but authentication failed")
	} else {
		t.Log("User authenticated successfully")
	}

	// Clean up
	cleanupUser(t, db, username)
}

func userExists(t *testing.T, db *sql.DB, username string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", username).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to check user existence: %v", err)
	}
	return count > 0
}

func cleanupUser(t *testing.T, db *sql.DB, username string) {
	t.Log("Cleaning up test user...")
	_, err := db.Exec("DELETE FROM users WHERE username = $1", username)
	if err != nil {
		t.Fatalf("Failed to clean up test user: %v", err)
	}
}
