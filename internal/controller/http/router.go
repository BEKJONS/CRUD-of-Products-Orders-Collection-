package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "ulab3/docs"
	"ulab3/internal/controller"
)

// title Api For ulab
// version 1.0
// description Ulab3 API
func NewRouter(engine *gin.Engine, ctr *controller.Controller) {

	// Use CORS middleware

	// Swagger documentation route
	engine.GET("/swagger/*eny", ginSwagger.WrapHandler(swaggerFiles.Handler))
	hp := NewProductHandler(ctr.Product)
	ho := NewOrderHandler(ctr.Order)
	// Define route groups
	products := engine.Group("/products")
	orders := engine.Group("/orders")

	// Define product routes
	products.POST("/", hp.CreateProduct)      // Create a new product
	products.GET("/", hp.GetAllProducts)      // Get all products
	products.GET("/:id", hp.GetProductByID)   // Get product by ID
	products.PUT("/:id", hp.UpdateProduct)    // Update a product
	products.DELETE("/:id", hp.DeleteProduct) // Delete a product

	// Define order routes
	orders.POST("/", ho.CreateOrder)      // Create a new order
	orders.GET("/", ho.GetAllOrders)      // Get all orders
	orders.GET("/:id", ho.GetOrderByID)   // Get order by ID
	orders.PUT("/:id", ho.UpdateOrder)    // Update an order
	orders.DELETE("/:id", ho.DeleteOrder) // Delete an order
}
