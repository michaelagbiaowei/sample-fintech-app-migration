package tests

import (
	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/auth"
	"github.com/michaelagbiaowei/sample-fintech-app-migration/internal/storage"
	"testing"
)

func TestRegisterAndAuthenticateUser(t *testing.T) {
	storage.InitDB()
	db := storage.GetDB()

	username := "testuser"
	password := "testpassword"

	err := auth.RegisterUser(db, username, password)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	authenticated, err := auth.AuthenticateUser(db, username, password)
	if err != nil {
		t.Fatalf("Error during authentication: %v", err)
	}

	if !authenticated {
		t.Errorf("Expected user to be authenticated")
	}

	// Clean up
	_, err = db.Exec("DELETE FROM users WHERE username = $1", username)
	if err != nil {
		t.Fatalf("Failed to clean up test user: %v", err)
	}
}