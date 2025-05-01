//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/santosjordi/posgoexpert/19-di/product"
)

var setRepository = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
