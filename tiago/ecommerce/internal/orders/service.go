package orders

import (
	"context"
	repo "ecommerce/internal/adapters/sqlc"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder CreateOrderParams) (repo.Order, error)
}

type Svc struct {
	repo *repo.Queries
	db *pgx.Conn
}

var (
	ErrProductNotFound = func (id int64) error {
		return fmt.Errorf("Product not found with productId: %s", id)
	}
	ErrProductNotInStock = func (id int64) error {
		return fmt.Errorf("Product with id %s has not enough quantity in stock", id)
	}
)

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &Svc { repo: repo, db: db }
}

func (this *Svc) PlaceOrder(ctx context.Context, tempOrder CreateOrderParams) (repo.Order, error) {
	if tempOrder.CustomerId == 0 {
		return repo.Order {}, fmt.Errorf("CustomerId is required.")
	}
	if len(tempOrder.Items) == 0 {
		return repo.Order {}, fmt.Errorf("At least 1 item is required to place an order.")
	}

	// Initialize transaction
	var tx, errTx = this.db.Begin(ctx)
	if errTx != nil {
		return repo.Order {}, fmt.Errorf("Error to begin transaction: %s", errTx)
	}
	defer tx.Rollback(ctx)
	var qtx = this.repo.WithTx(tx) // Transaction Query

	// Create order
	var order, errCreate = qtx.CreateOrder(ctx, tempOrder.CustomerId);
	if errCreate != nil {
		return repo.Order {}, errCreate
	}

	// Create order items
	for _, item := range tempOrder.Items {
		var prod, errProd = qtx.FindProductById(ctx, item.ProductId)
		if errProd != nil {
			return repo.Order {}, ErrProductNotFound(item.ProductId)
		}

		if item.Quantity > prod.Quantity.Int32 {
			return repo.Order {}, ErrProductNotInStock(item.ProductId)
		}

		var newOrderItem = repo.CreateOrderItemParams {
			OrderID: order.ID,
			ProductID: item.ProductId,
			Quantity: pgtype.Int4 { Int32: item.Quantity },
			PriceInCents: prod.PriceInCents,
		}
		var _, errCreate = qtx.CreateOrderItem(ctx, newOrderItem)
		if errCreate != nil {
			return repo.Order {}, errCreate
		}
	}

	return order, nil
}
