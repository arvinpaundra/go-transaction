package model

import (
	"time"

	"gorm.io/gorm"
)

// Define common table collumns
type Common struct {
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

// Create default data for created_at and updated_at
func (c *Common) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now
	return
}

// Create default data for updated_at if update happens to the row
func (c *Common) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}
