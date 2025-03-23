package database

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

// this is equivalent to a repository pattern

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) CreateCategory(name string, description string) (Category, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return Category{}, err
	}
	category := Category{ID: id, Name: name, Description: description}
	log.Printf("Created Category: %+v\n", category)
	return category, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	log.Printf("Retrieved Categories: %+v\n", categories)
	return categories, nil
}

// Given a courseID, find the category associated with it
func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow(`
        SELECT c.id, c.name, c.description 
        FROM categories c 
        JOIN courses co ON c.id = co.category_id 
        WHERE co.id = ?`, courseID).Scan(&id, &name, &description)

	if err != nil {
		log.Printf("Error finding category for courseID %s: %v", courseID, err)
		return Category{}, err
	}

	log.Printf("Found Category: ID=%s, Name=%s, Description=%s", id, name, description)
	return Category{ID: id, Name: name, Description: description}, nil
}
