package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Category has many Products, CategoryID is the foreign key
type Category struct {
	ID       int    `gorm:"primaryKey;autoIncrement;type:int(11)"`
	Name     string `gorm:"type:varchar(100)"`
	Products []Product
	gorm.Model
}

// Product has one SerialNumber, ProductID is the foreign key
type Product struct {
	ID           int      `gorm:"primaryKey;autoIncrement;type:int(11)"`
	Name         string   `gorm:"type:varchar(100)"`
	Price        float64  `gorm:"type:decimal(10,2)"`
	CategoryID   int      `gorm:"type:int(11)"`
	Category     Category `gorm:"foreignKey:CategoryID"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int    `gorm:"primaryKey;autoIncrement;type:int(11)"`
	Number    string `gorm:"type:varchar(14);unique"`
	ProductID int    `gorm:"type:int(11);unique"`
	gorm.Model
}

func main() {

	dsn := "root:root@tcp(localhost:3309)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Habilita suporte a relações de chave estrangeira
		PrepareStmt:          true,
		FullSaveAssociations: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	if err != nil {
		return
	}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err.Error())
	}

	for _, category := range categories {
		fmt.Printf("Categoria: %v \n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("	- Produto: %v | Serial: %v \n", product.Name, product.SerialNumber.Number)
		}
	}

	/*

		var categories []Category
		err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

		if err != nil {
			panic(err)
		}

		for _, category := range categories {
			println(category.Name, ":")
			for _, product := range category.Products {
				println(" - ", product.Name)
			}
		}*/

	/*
		SerialNumber := SerialNumber{
			ProductID: 5,
			Number:    "12345678901234",
		}

		db.Create(&SerialNumber)*/

	/*var category Category
	category.Name = "Category 1"
	db.Save(&category)

	var product Product
	product.Name = "Product 1"
	product.Price = 10.0
	product.CategoryID = category.ID
	db.Save(&product)*/

	// find product with id 1 and deleted_at

	/*
				    // create product
					products := []Product{
						{Name: "Product 1", Price: 10.0},
						{Name: "Product 2", Price: 20.0},
						{Name: "Product 3", Price: 30.0},
						{Name: "Product 4", Price: 40.0},
					}

					db.Create(&products)

					// find product with id 1
					var product Product
					db.First(&product, 1)

					// registros com deletd_at not null
					db.Unscoped().First(&product, 1)

					fmt.Printf("Product: %+v\n", product.Name)

					// find all products
					var products []Product
					db.Find(&products)

					for _, p := range products {
						fmt.Printf("Product: %+v\n", p.Name)
					}


			// delete product with id 1
			db.Where("id = ?", 1).Delete(&Product{})

			// update product with id 2
			db.Model(&Product{}).Where("id = ?", 2).Update("name", "Product 1 Updated")

			// filter
			var products []Product
			searchTerm := "Product 1"

			db.Where("name LIKE ?", "%"+searchTerm+"%").Find(&products)

			jsonData, err := json.MarshalIndent(products, "", "  ")
			if err != nil {
				panic(err)
			}

			fmt.Print(string(jsonData))

		var product Product
		db.Find(&product, 10)

		product.Name = "Product 10 Updated"
		db.Save(&product)

		fmt.Printf("Product: %+v\n", product.Name)

	*/

}
