package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

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

	// Seed data for the courses table
	courseSchema := `
	CREATE TABLE courses (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		category_id TEXT,
		FOREIGN KEY (category_id) REFERENCES categories(id)
	);`
	_, err = s.db.Exec(courseSchema)
	s.Require().NoError(err)

	// Seed data for courses
	courseSeedData := []struct {
		id, name, description, categoryID string
	}{
		{"1", "Gardening Basics", "Introduction to gardening", "1"},
		{"2", "Advanced Gardening", "In-depth gardening techniques", "1"},
		{"3", "Creative Writing", "Enhance your writing skills", "2"},
		{"4", "Fashion Design 101", "Basics of fashion design", "4"},
	}

	for _, seed := range courseSeedData {
		_, err := s.db.Exec(`INSERT INTO courses (id, name, description, category_id) VALUES (?, ?, ?, ?)`,
			seed.id, seed.name, seed.description, seed.categoryID)
		s.Require().NoError(err)
	}
}

func (s *CategoryTestSuite) TearDownTest() {
	_, err := s.db.Exec("DROP TABLE IF EXISTS courses")
	s.Require().NoError(err)
	_, err = s.db.Exec("DROP TABLE IF EXISTS categories")
	s.Require().NoError(err)
}

func (s *CategoryTestSuite) TearDownSuite() {
	s.db.Close()
}

func TestCategorySuite(t *testing.T) {
	suite.Run(t, new(CategoryTestSuite))
}

// Tests need to be assigned to the CategoryTestSuite and can't have any arguments passed in
func (s *CategoryTestSuite) TestCreateCategory_withValidParams_returnOK() {
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

// Test if the FindAll method returns all categories from the database
func (s *CategoryTestSuite) TestFindAll_withData_returnAllCategories() {
	// Arrange
	categoryDB := NewCategory(s.db)

	// Act
	categories, err := categoryDB.FindAll()

	// Assert
	s.NoError(err)
	s.Len(categories, 3)
}

// Test if the method FindByCourseID returns the correct category
func (s *CategoryTestSuite) TestFindByCourseID_withValidID_returnCategory() {
	// Arrange
	categoryDB := NewCategory(s.db)
	courseID := "1"

	// Act
	category, err := categoryDB.FindByCourseID(courseID)

	// Assert
	s.NoError(err)
	s.Equal("Gardening", category.Name)
	s.Equal("Tools and supplies for gardening", category.Description)
}
