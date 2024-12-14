package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ulab3/internal/entity"
	"ulab3/internal/usecase"
)

// OrderHandler handles HTTP requests for orders.
type OrderHandler struct {
	orderService *usecase.OrderService
}

// NewOrderHandler creates a new OrderHandler.
func NewOrderHandler(orderService *usecase.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order in the system.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body entity.Order true "Order data"
// @Success 201 {object} entity.Order
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order entity.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, entity.Error{Message: fmt.Sprintf("invalid request body: %v", err)})
		return
	}

	createdOrder, err := h.orderService.CreateOrder(c, &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Error{Message: fmt.Sprintf("failed to create order: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, createdOrder)
}

// GetAllOrders godoc
// @Summary Get all orders
// @Description Retrieve all orders from the system.
// @Tags orders
// @Produce  json
// @Success 200 {array} entity.Order
// @Failure 500 {object} entity.Error
// @Router /orders [get]
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.orderService.GetAllOrders(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Error{Message: fmt.Sprintf("failed to fetch orders: %v", err)})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrderByID godoc
// @Summary Get an order by ID
// @Description Retrieve an order by its ID.
// @Tags orders
// @Produce  json
// @Param id path string true "Order ID"
// @Success 200 {object} entity.Order
// @Failure 404 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	order, err := h.orderService.GetOrderByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, entity.Error{Message: fmt.Sprintf("order not found: %v", err)})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrder godoc
// @Summary Update an order
// @Description Update an existing order's details.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Param order body entity.Order true "Updated order data"
// @Success 200 {object} entity.Order
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order entity.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, entity.Error{Message: fmt.Sprintf("invalid request body: %v", err)})
		return
	}

	err := h.orderService.UpdateOrder(c, id, &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Error{Message: fmt.Sprintf("failed to update order: %v", err)})
		return
	}

	c.JSON(http.StatusOK, order)
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order from the system.
// @Tags orders
// @Param id path string true "Order ID"
// @Success 200 {object} entity.Order
// @Failure 500 {object} entity.Error
// @Router /orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := h.orderService.DeleteOrder(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Error{Message: fmt.Sprintf("failed to delete order: %v", err)})
		return
	}

	c.JSON(http.StatusOK, entity.Order{ID: id})
}
