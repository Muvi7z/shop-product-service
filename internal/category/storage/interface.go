package storage

import (
	"context"
	"shop-product-service/internal/category/modal"
)

type Storage interface {
	FindOne(ctx context.Context, id int64) (modal.Category, error)
	FindByCategory(ctx context.Context, categoryId int64) ([]modal.Category, error)
	Create(ctx context.Context, product modal.Category) (int64, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, product modal.Category) error
}
