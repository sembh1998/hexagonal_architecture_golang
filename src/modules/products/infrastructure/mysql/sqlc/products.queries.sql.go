// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: products.queries.sql

package productsqlc

import (
	"context"
)

const findAll = `-- name: FindAll :many
SELECT id, name, description, price FROM product
`

func (q *Queries) FindAll(ctx context.Context) ([]Product, error) {
	rows, err := q.query(ctx, q.findAllStmt, findAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
