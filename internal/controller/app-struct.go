package controller

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"ulab3/internal/usecase"
	"ulab3/internal/usecase/repo"
)

type Controller struct {
	Order   *usecase.OrderService
	Product *usecase.ProductService
}

func NewController(db *mongo.Client, log *slog.Logger, databaseName string) *Controller {
	// Get a reference to the collections
	productCollection := db.Database(databaseName).Collection("products")
	orderCollection := db.Database(databaseName).Collection("orders")

	// Initialize repositories
	productRepo := repo.NewProductRepository(productCollection)
	orderRepo := repo.NewOrderRepository(orderCollection)

	// Initialize services
	productService := usecase.NewProductService(productRepo, log)
	orderService := usecase.NewOrderService(orderRepo, productRepo, log)

	// Create and return the Controller instance
	return &Controller{
		Product: productService,
		Order:   orderService,
	}
}
