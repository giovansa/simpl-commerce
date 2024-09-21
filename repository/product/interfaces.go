package product

import (
	"context"
	"simpl-commerce/model/product"
)

type RepositoryInterface interface {
	FetchProducts(ctx context.Context, input FetchProductReq) ([]product.Product, error)
	CreateLedger(ctx context.Context, input NewLedger) error
}
