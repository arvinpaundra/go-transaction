package repository

import (
	"gorm.io/gorm"
)

// Define the interface for method call
type {{.EntityName}} interface {
}

// Define the scoped type
type {{.StructName}} struct {
	Db *gorm.DB
}

// Here the function is to get the database connection and to allow running the query
// This function will be called inside factory package
// gorm.DB contain database connection
func New{{.EntityName}}Repository(db *gorm.DB) *{{.StructName}} {
	return &{{.StructName}}{
		db,
	}
}
