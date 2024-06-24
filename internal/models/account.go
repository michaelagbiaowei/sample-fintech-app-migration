package models

import (
	"database/sql"
	"errors"
)

type Account struct {
	ID      int
	UserID  int
	Balance float64
}

func GetAccountByUsername(db *sql.DB, username string) (*Account, error) {
	account := &Account{}
	err := db.QueryRow(`
		SELECT a.id, a.user_id, a.balance 
		FROM accounts a 
		JOIN users u ON a.user_id = u.id 
		WHERE u.username = $1
	`, username).Scan(&account.ID, &account.UserID, &account.Balance)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *Account) Deposit(db *sql.DB, amount float64) error {
	_, err := db.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, a.ID)
	if err != nil {
		return err
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(db *sql.DB, amount float64) error {
	if a.Balance < amount {
		return errors.New("insufficient funds")
	}
	_, err := db.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, a.ID)
	if err != nil {
		return err
	}
	a.Balance -= amount
	return nil
}