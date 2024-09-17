package storage

import (
	"context"
	"shop-product-service/internal/product/modal"
)

type Storage interface {
	FindOne(ctx context.Context, id string) (modal.Product, error)
	FindByCategory(ctx context.Context, categoryId int64) ([]modal.Product, error)
	Create(ctx context.Context, product modal.Product) (string, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, product modal.Product) error
}
