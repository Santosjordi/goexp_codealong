package main

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/santosjordi/posgoexpert/17-sqlc/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	err = queries.CreateCategory(
		ctx, db.CreateCategoryParams{
			ID:          uuid.New().String(),
			Name:        "Test",
			Description: sql.NullString{String: "Test", Valid: true},
		})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
