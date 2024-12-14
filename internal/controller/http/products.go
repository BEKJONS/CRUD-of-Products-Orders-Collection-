package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ulab3/internal/entity"
	"ulab3/internal/usecase"
)

// ProductHandler handles HTTP requests for products.
type ProductHandler struct {
	productService *usecase.ProductService
}

// NewProductHandler creates a new ProductHandler.
func NewProductHandler(productService *usecase.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product in the system.
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body entity.Product true "Product data"
// @Success 201 {object} entity.Product
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, entity.Error{Message: fmt.Sprintf("invalid request body: %v", err)})
		return
	}

	createdProduct, err := h.productService.CreateProduct(c, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Error{Message: fmt.Sprintf("failed to create product: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve all products from the system.
// @Tags products
// @Produce  json
// @Success 200 {array} entity.Product
// @Failure 500 {object} entity.Error
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Error{Message: fmt.Sprintf("failed to fetch products: %v", err)})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Retrieve a product by its ID.
// @Tags products
// @Produce  json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 404 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productService.GetProductByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, entity.Error{Message: fmt.Sprintf("product not found: %v", err)})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update an existing product's details.
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Param product body entity.Product true "Updated product data"
// @Success 200 {object} entity.Product
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, entity.Error{Message: fmt.Sprintf("invalid request body: %v", err)})
		return
	}

	err := h.productService.UpdateProduct(c, id, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Error{Message: fmt.Sprintf("failed to update product: %v", err)})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product from the system.
// @Tags products
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 500 {object} entity.Error
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := h.productService.DeleteProduct(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Error{Message: fmt.Sprintf("failed to delete product: %v", err)})
		return
	}

	c.JSON(http.StatusOK, entity.Product{ID: id})
}
