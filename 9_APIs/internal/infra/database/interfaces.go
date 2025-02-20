package database

import "github.com/santosjordi/posgoexp/9_apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type ProdcutInterface interface {
	Create(product *entity.Product) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(productID string) (*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	Delete(productID string) error
}
