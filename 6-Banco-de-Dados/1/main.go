package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, p *Product) error {

	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	_, err = stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("update products set name=?, price=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select * from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select * from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil

}

func deleteProduct(db *sql.DB, id string) error {
	_, err := db.Exec("delete from products where id =?", id)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/goexpert")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = deleteProduct(db, "45751634-39ee-49e0-b5ed-7f6032d9f13d")
	if err != nil {
		panic(err)
	}

	/*products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("O produto %v tem o valor de %.2f \n", p.Name, p.Price)
	}*/

	/*
		product := NewProduct("Clean Architecture", 59.00)
		err = insertProduct(db, product)
		if err != nil {
			panic(err)
		}

		product.Price = 49.00
		err = updateProduct(db, product)
		if err != nil {
			panic(err)
		}

		p, err := selectProduct(db, product.ID)
		if err != nil {
			panic(err)
		}

		fmt.Printf("O produto %v tem o valor de %.2f", p.Name, p.Price)
	*/

}
