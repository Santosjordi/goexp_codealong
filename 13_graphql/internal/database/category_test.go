package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CategoryTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (s *CategoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)
	s.db = db
}

func (s *CategoryTestSuite) SetupTest() {
	schema := `
    CREATE TABLE categories (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT
    );`
	_, err := s.db.Exec(schema)
	s.Require().NoError(err)

	// Seed data
	seedData := []struct {
		id, name, description string
	}{
		{"1", "Gardening", "Tools and supplies for gardening"},
		{"2", "Writing", "Lero-Lero generator"},
		{"4", "Fashion Design", "Apparel and accessories"},
	}

	for _, seed := range seedData {
		_, err := s.db.Exec(`INSERT INTO categories (id, name, description) VALUES (?, ?, ?)`,
			seed.id, seed.name, seed.description)
		s.Require().NoError(err)
	}
}

func (s *CategoryTestSuite) TearDownTest() {
	_, err := s.db.Exec("DROP TABLE IF EXISTS categories")
	s.Require().NoError(err)
}

func (s *CategoryTestSuite) TearDownSuite() {
	s.db.Close()
}

func TestCategorySuite(t *testing.T) {
	suite.Run(t, new(CategoryTestSuite))
}

func (s *CategoryTestSuite) TestCreateCategory_withValidParams_returnOK(t *testing.T) {
	// Arrange
	categoryDB := NewCategory(s.db)
	name := "Technology"
	description := "Tech related content"

	// Act
	category, err := categoryDB.CreateCategory(name, description)

	// Assert
	s.NoError(err)
	s.NotEmpty(category.ID)
	s.Equal(name, category.Name)
	s.Equal(description, category.Description)

	// Verify in database
	var savedCategory Category
	err = s.db.QueryRow("SELECT id, name, description FROM categories WHERE id = ?",
		category.ID).Scan(&savedCategory.ID, &savedCategory.Name, &savedCategory.Description)

	s.NoError(err)
	s.Equal(category.ID, savedCategory.ID)
	s.Equal(name, savedCategory.Name)
	s.Equal(description, savedCategory.Description)
}
