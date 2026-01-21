package models

import (
	"time"
)

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
	Stock     int       `json:"stock" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Request structs for input validation
type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required,min=1,max=100"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Stock int     `json:"stock" binding:"gte=0"`
}

type UpdateProductRequest struct {
	Name  *string  `json:"name,omitempty" binding:"omitempty,min=1,max=100"`
	Price *float64 `json:"price,omitempty" binding:"omitempty,gt=0"`
	Stock *int     `json:"stock,omitempty" binding:"omitempty,gte=0"`
}
