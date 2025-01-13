package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// ======= HasMany
	// var categories []Category
	// err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	// if err != nil {
	// 	panic(err)
	// }
	// for _, category := range categories {
	// 	fmt.Println(category.CategoryName, ":")
	// 	for _, product := range category.Products {
	// 		println("-", product.ProductName)
	// 	}
	// }

	// ======= Has One
	// category := Category{
	// 	CategoryName: "Eletronics",
	// 	Description:  "Eletronic products",
	// 	Picture:      "",
	// 	CreatedAt:    time.Now(),
	// 	UpdatedAt:    time.Now(),
	// }
	// db.Create(&category)

	// product := Product{
	// 	ProductName: "Smartphone",
	// 	CategoryID:  13,
	// }
	// db.Create(&product)

	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 77,
	// })

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	println(product.ProductName, product.Category.CategoryName, product.SerialNumber.Number)
	// }
	// ===== Has Many
	// var categories []Category
	// // Pegadinha do has Many, carregar os serial numbers a partir da categoria
	// err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	// if err != nil {
	// 	panic(err)
	// }
	// for _, category := range categories {
	// 	fmt.Println(category.CategoryName, ":")
	// 	for _, product := range category.Products {
	// 		println("-", product.ProductName, "[", product.SerialNumber.Number, "]")
	// 	}
	// }

}
