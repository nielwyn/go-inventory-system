package models

import (
	"time"

	"gorm.io/gorm"
)

// Item represents an inventory item
type Item struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	SKU         string         `gorm:"uniqueIndex;not null" json:"sku"`
	Description string         `json:"description"`
	Quantity    int            `gorm:"not null;default:0" json:"quantity"`
	Price       float64        `gorm:"not null;default:0" json:"price"`
	Category    string         `json:"category"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Item
func (Item) TableName() string {
	return "items"
}

// CreateItemRequest represents a request to create an item
type CreateItemRequest struct {
	Name        string  `json:"name" binding:"required,min=1,max=200"`
	SKU         string  `json:"sku" binding:"required,min=1,max=100"`
	Description string  `json:"description" binding:"max=1000"`
	Quantity    int     `json:"quantity" binding:"non_negative"`
	Price       float64 `json:"price" binding:"non_negative"`
	Category    string  `json:"category" binding:"max=100"`
}

// UpdateItemRequest represents a request to update an item
type UpdateItemRequest struct {
	Name        *string  `json:"name" binding:"omitempty,min=1,max=200"`
	SKU         *string  `json:"sku" binding:"omitempty,min=1,max=100"`
	Description *string  `json:"description" binding:"omitempty,max=1000"`
	Quantity    *int     `json:"quantity" binding:"omitempty,non_negative"`
	Price       *float64 `json:"price" binding:"omitempty,non_negative"`
	Category    *string  `json:"category" binding:"omitempty,max=100"`
}
