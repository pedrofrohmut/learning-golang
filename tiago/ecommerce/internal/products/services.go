package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) error
}

type Svc struct {}

func NewSvc() *Svc {
	return &Svc {}
}

func (this *Svc) ListProducts(ctx context.Context) error {
	return nil
}
