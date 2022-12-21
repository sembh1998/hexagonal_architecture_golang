-- name: FindAll :many
SELECT * FROM product;

-- name: Save :execresult
INSERT INTO product(name, description, price) VALUES(sqlc.arg(name), sqlc.arg(description), sqlc.narg(price));