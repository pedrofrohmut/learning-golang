-- name: ListProducts :many
select * from products;

-- name: FindProductById :one
select * from products where id = $1;

-- name: CreateOrder :one
insert into orders (customer_id) values ($1) returning *;

-- name: CreateOrderItem :one
insert into order_items (product_id, quantity, price_in_cents, order_id) values ($1, $2, $3, $4) returning *;
