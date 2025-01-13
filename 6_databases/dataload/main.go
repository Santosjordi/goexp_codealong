package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ProductID       int      `gorm:"primaryKey;autoIncrement" csv:"product_id"`
	ProductName     string   `csv:"product_name"`
	SupplierID      int      `csv:"supplier_id"`
	QuantityPerUnit string   `csv:"quantity_per_unit"`
	UnitPrice       float64  `csv:"unit_price"`
	UnitsInStock    int      `csv:"units_in_stock"`
	UnitsOnOrder    int      `csv:"units_on_order"`
	ReorderLevel    int      `csv:"reorder_level"`
	Discontinued    bool     `csv:"discontinued"`
	CategoryID      int      `csv:"category_id"`
	Category        Category `gorm:"foreignKey:CategoryID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	CategoryID   int    `gorm:"primaryKey;autoIncrement" csv:"category_id"`
	CategoryName string `csv:"category_name"`
	Description  string `csv:"description"`
	Picture      string `csv:"picture"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// Carrega os dados dos CVS do northwind traders para o banco mysql

func main() {
	// Connect to the database
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Drop the table if it exists
	err = db.Migrator().DropTable(&Product{}, &Category{})
	if err != nil {
		log.Fatalf("Failed to drop table: %v", err)
	}

	// AutoMigrate tables
	db.AutoMigrate(&Product{}, &Category{})

	// Load Products
	err = loadCSVToDB("/home/jordi/goexpert_pos/6_databases/products.csv", Product{}, db)
	if err != nil {
		fmt.Println("Error loading products:", err)
		return
	}

	// Load Categories
	err = loadCSVToDB("/home/jordi/goexpert_pos/6_databases/categories.csv", Category{}, db)
	if err != nil {
		fmt.Println("Error loading categories:", err)
		return
	}

	fmt.Println("Data loaded successfully")
}

func loadCSVToDB(filepath string, model interface{}, db *gorm.DB) error {
	// Open the CSV file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the CSV content
	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	if len(records) < 2 {
		return fmt.Errorf("CSV file must have at least one data row")
	}

	// Extract the headers and data rows
	headers := records[0]
	rows := records[1:]

	// Reflect on the model's type
	modelType := reflect.TypeOf(model)
	if modelType.Kind() != reflect.Struct {
		return fmt.Errorf("model must be a struct")
	}

	// Map each row to the model and save to DB
	for _, row := range rows {
		newModel := reflect.New(modelType).Interface()

		// Map the CSV row to the struct
		err := mapCSVRowToStruct(headers, row, newModel)
		if err != nil {
			return fmt.Errorf("failed to map CSV row: %w", err)
		}

		// Insert the data into the database
		if err := db.Create(newModel).Error; err != nil {
			return fmt.Errorf("failed to insert data into database: %w", err)
		}
	}

	return nil
}

func mapCSVRowToStruct(headers []string, row []string, model interface{}) error {
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	for i, header := range headers {
		field := modelValue.FieldByNameFunc(func(name string) bool {
			field, _ := modelType.FieldByName(name)
			return field.Tag.Get("csv") == header
		})
		if !field.IsValid() {
			continue
		}

		// Convert and set the field value
		err := setFieldValue(field, row[i])
		if err != nil {
			return fmt.Errorf("failed to set field %s: %w", header, err)
		}
	}
	return nil
}

func setFieldValue(field reflect.Value, value string) error {
	if !field.CanSet() {
		return fmt.Errorf("field is not settable")
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(intValue))
	case reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(floatValue)
	case reflect.Bool:
		boolValue := value == "1" || value == "true"
		field.SetBool(boolValue)
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}
	return nil
}
