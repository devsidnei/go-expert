package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Category has many Products, CategoryID is the foreign key
type Category struct {
	ID       int       `gorm:"primaryKey;autoIncrement;type:int(11)"`
	Name     string    `gorm:"type:varchar(100)"`
	Products []Product `gorm:"many2many:products_categories"`
	gorm.Model
}

// Product has one SerialNumber, ProductID is the foreign key
type Product struct {
	ID         int        `gorm:"primaryKey;autoIncrement;type:int(11)"`
	Name       string     `gorm:"type:varchar(100)"`
	Price      float64    `gorm:"type:decimal(10,2)"`
	Categories []Category `gorm:"many2many:products_categories"`
	gorm.Model
}

func main() {

	dsn := "root:root@tcp(localhost:3309)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:          true,
		FullSaveAssociations: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Product{}, &Category{})
	if err != nil {
		return
	}

	/*categories := []Category{
		{Name: "Categoria 1"},
		{Name: "Categoria 2"},
		{Name: "Categoria 3"},
	}
	db.Create(&categories)*/

	var categories []Category
	db.Find(&categories)

	product := Product{
		Name:       "Produto 1",
		Categories: categories,
	}

	db.Create(&product)

	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err.Error())
	}

	for _, category := range categories {
		fmt.Printf("Categoria: %v \n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("	- Produto: %v \n", product.Name)
		}
	}
}
