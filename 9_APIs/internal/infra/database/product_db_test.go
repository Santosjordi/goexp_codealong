package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/santosjordi/posgoexp/9_apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	// arrange
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	// act
	product, err := entity.NewProduct("test", 10.0)
	productDB := NewProduct(db)

	// assert
	assert.Nil(t, err)
	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ProductID)

}

func TestFindAllProducts(t *testing.T) {
	// arrange
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	for i := 10; i < 34; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)

	// act
	products, err := productDB.FindAll(1, 10, "asc")

	// assert
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 10", products[0].Name)
	assert.Equal(t, "Product 19", products[9].Name)

	// act
	products_pg_two, err := productDB.FindAll(2, 10, "asc")

	// assert which elements are in the second page
	assert.Nil(t, err)
	assert.Len(t, products_pg_two, 10)
	assert.Equal(t, "Product 20", products_pg_two[0].Name)
	assert.Equal(t, "Product 29", products_pg_two[9].Name)

	// act
	products_pg_three, err := productDB.FindAll(3, 10, "asc")

	// assert last page has fewer elements
	assert.Nil(t, err)
	assert.Len(t, products_pg_three, 4)
	assert.Equal(t, "Product 30", products_pg_three[0].Name)
	assert.Equal(t, "Product 33", products_pg_three[3].Name)

}

func TestFindProductByID(t *testing.T) {
	// arrange
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	db.Create(product)

	productDB := NewProduct(db)

	// act
	productFound, err := productDB.FindByID(product.ProductID.String())

	// assert
	assert.Nil(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	// arrange
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	db.Create(product)

	productDB := NewProduct(db)

	// assert that there is only one product in the database
	var count int64
	db.Model(&entity.Product{}).Count(&count)
	assert.Equal(t, int64(1), count)

	// act
	product.Name = "Product 2"
	product.Price = 20.0
	err = productDB.Update(product)
	assert.Nil(t, err)
	productFound, err := productDB.FindByID(product.ProductID.String())

	// assert
	assert.Nil(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	// assert that there is still only one product in the database
	db.Model(&entity.Product{}).Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestDeleteProduct(t *testing.T) {
	// arrange
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	db.Create(product)

	productDB := NewProduct(db)

	// assert that there is only one product in the database
	var count int64
	db.Model(&entity.Product{}).Count(&count)
	assert.Equal(t, int64(1), count)

	// act
	err = productDB.Delete(product.ProductID.String())

	// assert
	assert.Nil(t, err)
	// assert that there are no products in the database
	db.Model(&entity.Product{}).Count(&count)
	assert.Equal(t, int64(0), count)
}
