// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"database/sql"
)

type Category struct {
	ID          string
	Name        string
	Description sql.NullString
}

type Course struct {
	ID          string
	Name        string
	Description sql.NullString
	CategoryID  sql.NullString
	Price       sql.NullString
}
