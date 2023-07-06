// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     []string  `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64         `json:"id"`
	AccountID sql.NullInt64 `json:"account_id"`
	// only positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID           int64         `json:"id"`
	FromAccoutID sql.NullInt64 `json:"from_accout_id"`
	ToAccountID  sql.NullInt64 `json:"to_account_id"`
	// only positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
