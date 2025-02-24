package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/santosjordi/posgoexp/9_apis/configs"
	"github.com/santosjordi/posgoexp/9_apis/internal/entity"
	"github.com/santosjordi/posgoexp/9_apis/internal/infra/database"
	"github.com/santosjordi/posgoexp/9_apis/internal/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	// repository receives the db instance
	productDB := database.NewProduct(db)

	// handler receives the repo, the repo needs to have all interface methods implemented
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetAllProducts)

	http.ListenAndServe(config.WebServerPort, r)
}
