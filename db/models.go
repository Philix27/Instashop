// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Orderstatus string

const (
	OrderstatusCREATED   Orderstatus = "CREATED"
	OrderstatusCOMPLETED Orderstatus = "COMPLETED"
	OrderstatusAPPEAL    Orderstatus = "APPEAL"
)

func (e *Orderstatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Orderstatus(s)
	case string:
		*e = Orderstatus(s)
	default:
		return fmt.Errorf("unsupported scan type for Orderstatus: %T", src)
	}
	return nil
}

type NullOrderstatus struct {
	Orderstatus Orderstatus
	Valid       bool // Valid is true if Orderstatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderstatus) Scan(value interface{}) error {
	if value == nil {
		ns.Orderstatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Orderstatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderstatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Orderstatus), nil
}

type Order struct {
	ID          int64
	UserID      pgtype.Int4
	OrderStatus NullOrderstatus
	Items       []interface{}
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Orderitem struct {
	ID        int64
	ProductID pgtype.Int4
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Product struct {
	ID          int64
	Title       string
	Description string
	ImageUrl    string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type User struct {
	ID             int64
	Wallets        []string
	FirstName      pgtype.Text
	LastName       pgtype.Text
	MiddleName     pgtype.Text
	Dob            pgtype.Timestamp
	Email          string
	Phone          pgtype.Text
	HashedPassword string
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}
