package main

import (
	"database/sql"
	"fmt"

	"github.com/santosjordi/posgoexpert/19-di/product"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	repository := product.NewProductRepository(db)

	usercase := product.NewProductUseCase(repository)

	product, err := usercase.GetProductByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
