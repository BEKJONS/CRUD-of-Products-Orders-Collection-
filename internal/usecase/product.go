package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"ulab3/internal/entity"
)

type ProductService struct {
	productRepo ProductRepository
	logger      *slog.Logger
}

func NewProductService(productRepo ProductRepository, logger *slog.Logger) *ProductService {
	return &ProductService{
		productRepo: productRepo,
		logger:      logger,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	s.logger.Info("Creating product:", product.Name)

	existingProduct, err := s.productRepo.FindByID(ctx, product.ID)
	if err == nil && existingProduct != nil {
		s.logger.Info("Product already exists:", product.ID)
		return nil, fmt.Errorf("product already exists")
	}

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	createdProduct, err := s.productRepo.Create(ctx, product)
	if err != nil {
		s.logger.Error("Failed to create product:", err)
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	s.logger.Info("Product created successfully:", createdProduct.ID)
	return createdProduct, nil
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	s.logger.Info("Fetching all products")

	products, err := s.productRepo.FindAll(ctx)
	if err != nil {
		s.logger.Error("Failed to fetch products:", err)
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	return products, nil
}

func (s *ProductService) GetProductByID(ctx context.Context, id string) (*entity.Product, error) {
	s.logger.Info("Fetching product by ID:", id)

	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		s.logger.Error("Product not found:", err)
		return nil, fmt.Errorf("product not found: %w", err)
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id string, product *entity.Product) error {
	s.logger.Info("Updating product:", id)

	product.UpdatedAt = time.Now()
	err := s.productRepo.Update(ctx, id, product)
	if err != nil {
		s.logger.Error("Failed to update product:", err)
		return fmt.Errorf("failed to update product: %w", err)
	}

	s.logger.Info("Product updated successfully:", id)
	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	s.logger.Info("Deleting product:", id)

	err := s.productRepo.Delete(ctx, id)
	if err != nil {
		s.logger.Error("Failed to delete product:", err)
		return fmt.Errorf("failed to delete product: %w", err)
	}

	s.logger.Info("Product deleted successfully:", id)
	return nil
}
