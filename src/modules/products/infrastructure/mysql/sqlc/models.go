// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package productsqlc

import (
	"database/sql"
)

type Product struct {
	ID          int32
	Name        string
	Description string
	Price       sql.NullFloat64
}
