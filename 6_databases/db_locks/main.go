package main

// Demonstração do uso de lock pessimista e otimista em operações com banco de dados

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	ProductID       int      `gorm:"primaryKey;autoIncrement" csv:"product_id"`
	ProductName     string   `csv:"product_name"`
	SupplierID      int      `csv:"supplier_id"`
	QuantityPerUnit string   `csv:"quantity_per_unit"`
	UnitPrice       float64  `csv:"unit_price"`
	UnitsInStock    int      `csv:"units_in_stock"`
	UnitsOnOrder    int      `csv:"units_on_order"`
	ReorderLevel    int      `csv:"reorder_level"`
	Discontinued    bool     `csv:"discontinued"`
	CategoryID      int      `csv:"category_id"`
	Category        Category `gorm:"foreignKey:CategoryID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	SerialNumber SerialNumber
}

type Category struct {
	CategoryID   int    `gorm:"primaryKey;autoIncrement" csv:"category_id"`
	CategoryName string `csv:"category_name"`
	Description  string `csv:"description"`
	Picture      string `csv:"picture"`
	Products     []Product
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	// Lock pessimista
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.CategoryName = "Bebidas"
	tx.Debug().Save(&c)
	tx.Commit()
}
