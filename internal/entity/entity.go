package entity

import "time"

type Product struct {
	ID        string    `json:"id" bson:"id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Price     float64   `json:"price" bson:"price"`
	Stock     int       `json:"stock" bson:"stock"`
	Category  string    `json:"category" bson:"category"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
type Order struct {
	ID         string    `json:"id" bson:"id,omitempty"`
	ProductID  string    `json:"product_id" bson:"product_id"`
	Quantity   int       `json:"quantity" bson:"quantity"`
	TotalPrice float64   `json:"total_price" bson:"total_price"`
	Status     string    `json:"status" bson:"status"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}
type Error struct {
	Message string `json:"message"`
}
