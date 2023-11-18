package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error

	if err != nil {
		panic(err.Error())
	}

	c.Name = "Categoria Com Lock"
	tx.Debug().Save(&c)
	tx.Commit()

}
