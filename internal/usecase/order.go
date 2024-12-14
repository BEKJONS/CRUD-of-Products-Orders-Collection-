package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"ulab3/internal/entity"
)

type OrderService struct {
	orderRepo   OrderRepository
	productRepo ProductRepository
	logger      *slog.Logger
}

func NewOrderService(orderRepo OrderRepository, productRepo ProductRepository, logger *slog.Logger) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
		logger:      logger,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	s.logger.Info("Creating order for product ID:", order.ProductID)

	// Check if product exists
	product, err := s.productRepo.FindByID(ctx, order.ProductID)
	if err != nil {
		s.logger.Error("Product not found:", err)
		return nil, fmt.Errorf("invalid product ID")
	}

	// Check stock availability
	if product.Stock < order.Quantity {
		s.logger.Info("Insufficient stock for product:", product.Name)
		return nil, fmt.Errorf("insufficient stock")
	}

	// Calculate the total price for the order
	order.TotalPrice = float64(order.Quantity) * product.Price
	order.Status = "Pending"
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	// Create the order
	createdOrder, err := s.orderRepo.Create(ctx, order)
	if err != nil {
		s.logger.Error("Failed to create order:", err)
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	// Update product stock
	product.Stock -= order.Quantity
	err = s.productRepo.Update(ctx, product.ID, product)
	if err != nil {
		s.logger.Error("Failed to update product stock:", err)
		return nil, fmt.Errorf("failed to update product stock: %w", err)
	}

	s.logger.Info("Order created successfully:", createdOrder.ID)
	return createdOrder, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]entity.Order, error) {
	s.logger.Info("Fetching all orders")

	orders, err := s.orderRepo.FindAll(ctx)
	if err != nil {
		s.logger.Error("Failed to fetch orders:", err)
		return nil, fmt.Errorf("failed to fetch orders: %w", err)
	}

	return orders, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, id string) (*entity.Order, error) {
	s.logger.Info("Fetching order by ID:", id)

	order, err := s.orderRepo.FindByID(ctx, id)
	if err != nil {
		s.logger.Error("Order not found:", err)
		return nil, fmt.Errorf("order not found: %w", err)
	}

	return order, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, id string, order *entity.Order) error {
	s.logger.Info("Updating order:", id)

	order.UpdatedAt = time.Now()
	err := s.orderRepo.Update(ctx, id, order)
	if err != nil {
		s.logger.Error("Failed to update order:", err)
		return fmt.Errorf("failed to update order: %w", err)
	}

	s.logger.Info("Order updated successfully:", id)
	return nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, id string) error {
	s.logger.Info("Deleting order:", id)

	err := s.orderRepo.Delete(ctx, id)
	if err != nil {
		s.logger.Error("Failed to delete order:", err)
		return fmt.Errorf("failed to delete order: %w", err)
	}

	s.logger.Info("Order deleted successfully:", id)
	return nil
}
