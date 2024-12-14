package usecase

import (
	"context"
	"ulab3/internal/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
	FindAll(ctx context.Context) ([]entity.Product, error)
	FindByID(ctx context.Context, id string) (*entity.Product, error)
	Update(ctx context.Context, id string, product *entity.Product) error
	Delete(ctx context.Context, id string) error
}

type OrderRepository interface {
	Create(ctx context.Context, order *entity.Order) (*entity.Order, error)
	FindAll(ctx context.Context) ([]entity.Order, error)
	FindByID(ctx context.Context, id string) (*entity.Order, error)
	Update(ctx context.Context, id string, order *entity.Order) error
	Delete(ctx context.Context, id string) error
}
