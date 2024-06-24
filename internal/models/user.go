package models

import (
	"database/sql"
)

type User struct {
	ID       int
	Username string
	Password string
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	user := &User{}
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateAccount(db *sql.DB, userID int) error {
	_, err := db.Exec("INSERT INTO accounts (user_id, balance) VALUES ($1, 0)", userID)
	return err
}