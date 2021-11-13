// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type BankTransaction struct {
	ID        int32
	Amount    sql.NullFloat64
	UserID    sql.NullInt32
	CreatedAt time.Time
}

type Batch struct {
	ID         int32
	Dispatched sql.NullBool
	Amount     float64
	UserID     sql.NullInt64
	CreatedAt  time.Time
}

type Transaction struct {
	ID        int32
	Amount    sql.NullFloat64
	UserID    sql.NullInt64
	CreatedAt time.Time
}

type User struct {
	ID        int32
	Name      sql.NullString
	CreatedAt time.Time
}
