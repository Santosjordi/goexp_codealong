package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/santosjordi/posgoexp/9_apis/internal/dto"
	"github.com/santosjordi/posgoexp/9_apis/internal/entity"
	"github.com/santosjordi/posgoexp/9_apis/internal/infra/database"
	entityPkg "github.com/santosjordi/posgoexp/9_apis/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product with the provided name and price
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateProductInput true "Product request"
// @Success      201
// @Failure      400 {object} Error
// @Router       /products [post]
// @Security	 ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary      Get a product by ID
// @Description  Retrieve a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path string true "Product ID" Format(uuid)
// @Success      200 {object} entity.Product
// @Failure      500 {object} Error
// @Failure      404 {object} Error
// @Router       /products/{id} [get]
// @Security	 ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary      Update a product by ID
// @Description  Update a product's details by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path string true "Product ID" Format(uuid)
// @Param        request body entity.Product true "Product details"
// @Success      200
// @Failure      404 {object} Error
// @Failure      500 {object} Error
// @Router       /products/{id} [put]
// @Security	 ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ProductID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary      Delete a product by ID
// @Description  Delete a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path string true "Product ID" Format(uuid)
// @Success      200
// @Failure      404 {object} Error
// @Failure      500 {object} Error
// @Router       /products/{id} [delete]
// @Security	 ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetProducts handles the HTTP request to retrieve a list of products.
// @Summary 	Get products
// @Description Get a list of products with pagination and sorting options
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		page 		query 		string 			false "Page number" 				default(0)
// @Param 		limit 		query 		string 			false "Number of items per page" 	default(10)
// @Param 		sort 		query 		string 			false "Sort order"
// @Success 	200 		{array} 	entity.Product
// @Failure 	404 		{object} 	Error
// @Failure 	500 		{object} 	Error
// @Router 		/products 	[get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
