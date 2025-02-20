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
