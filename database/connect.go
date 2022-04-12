// Package database provides primitives for database connections
package database

import (
	"github.com/scopehs/tutorial/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect opens a connection to the database
// Panics if cannot connect.
func Connect() {
	database, err := gorm.Open(mysql.Open("root:tut@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = database

	// Auto migrate User Model.
	database.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.MarkerPrices{},
	)
}
