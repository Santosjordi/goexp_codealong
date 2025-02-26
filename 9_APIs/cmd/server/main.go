package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
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
	userDB := database.NewUser(db)

	// handler receives the repo, the repo needs to have all interface methods implemented
	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JwtExpiresIn)

	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// custom middleware
	r.Use(LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)

		r.Put("/{id}", productHandler.UpdateProduct)

		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)

		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate-jwt", userHandler.GetJwt)
	r.Get("/users/{email}", userHandler.FindUserByEmail)

	http.ListenAndServe(config.WebServerPort, r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
