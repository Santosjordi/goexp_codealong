package product

type ProductUseCase struct {
	repo ProductRepositoryInterface
}

func NewProductUseCase(repo ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

// GetproductByID retrieves a product by its ID.
// This product was not supposed to be returned, we shoudl return a DTO instead,
// but for now we will return the entity
func (uc *ProductUseCase) GetProductByID(id int) (Product, error) {
	return uc.repo.GetProductByID(id)
}
