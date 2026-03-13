package products

import (
	"context"
	repo "ecommerce/internal/adapters/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
}

type Svc struct {
	repo repo.Querier
}

func NewSvc(repo repo.Querier) *Svc {
	return &Svc { repo: repo }
}

func (this *Svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return this.repo.ListProducts(ctx)
}
