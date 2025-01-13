package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Drop the table if it exists
	err = db.Migrator().DropTable(&Product{})
	if err != nil {
		log.Fatalf("Failed to drop table: %v", err)
	}
	// Create the table
	db.AutoMigrate(&Product{})

	// db.Create(&Product{Name: "Laptop", Price: 1000})

	// Create a batch of products
	products := []Product{
		{Name: "Laptop", Price: 1000},
		{Name: "Mouse", Price: 50},
		{Name: "Keyboard", Price: 150},
		{Name: "Monitor", Price: 300},
	}
	db.Create(&products)

	// Read one
	// var product Product
	// db.First(&product, "price = ?", "300")
	// fmt.Printf("Product: %+v\n", product.Name)

	// Read all
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Printf("Product: %+v\n", product.Name)
	// }

	//Where
	// var products []Product
	// db.Where("price > ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// Update
	var p Product
	db.First(&p, 1)
	p.Name = "New Laptop"
	db.Save(&p)

	// Delete
	var p2 Product
	db.First(&p2, 1)
	db.Delete(&p2)
}
