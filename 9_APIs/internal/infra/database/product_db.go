package database

import (
	"github.com/santosjordi/posgoexp/9_apis/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindByID(productID string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "product_id=?", productID).Error
	return &product, err
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	// Comportamento padrão nesse caso é retornar a ordenação em ordem crescente
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	// Se page e limit forem 0, então não há paginação
	if page != 0 && limit != 0 {
		err = p.DB.Offset((page - 1) * limit).Limit(limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ProductID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *Product) Delete(productID string) error {
	product, err := p.FindByID(productID)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
