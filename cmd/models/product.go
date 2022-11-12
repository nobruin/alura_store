package models

import (
	"alura_store/cmd/infra/database"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Quantity    int
	Description string
}

func NewModelProduct(name string, price float64, quantity int, description string) Product {
	return Product{
		Name:        name,
		Price:       price,
		Quantity:    quantity,
		Description: description,
	}
}

func NewModelCompleteProduct(id int, name string, price float64, quantity int, description string) Product {
	return Product{
		ID:          id,
		Name:        name,
		Price:       price,
		Quantity:    quantity,
		Description: description,
	}
}

func FindAll() []Product {
	var db = database.DatabaseConnect()
	defer db.Close()
	rows, err := db.Query("select id, name, description, price, quantity from products")

	if err != nil {
		panic(err.Error())
	}
	var products []Product
	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64
		err := rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p := NewModelCompleteProduct(id, name, price, quantity, description)
		products = append(products, p)
	}
	return products
}

func FindById(id string) Product {
	db := database.DatabaseConnect()
	rows, err := db.Query("select id, name, description, price, quantity from products where id = $1;", id)
	if err != nil {
		panic(err.Error())
	}
	var ID, quantity int
	var name, description string
	var price float64

	for rows.Next() {
		err := rows.Scan(&ID, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
	}
	p := NewModelCompleteProduct(ID, name, price, quantity, description)
	defer db.Close()
	return p
}

func CreateProduct(product Product) {
	db := database.DatabaseConnect()

	insert, err := db.Prepare("insert into products(name, price, quantity, description) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(product.Name, product.Price, product.Quantity, product.Description)
	defer db.Close()
}

func Delete(id int) {
	db := database.DatabaseConnect()

	delete, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}

func Update(product Product) {
	db := database.DatabaseConnect()
	insert, err := db.Prepare("update products set name=$1, price=$2, quantity=$3, description=$4 where id=$5;")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(product.Name, product.Price, product.Quantity, product.Description, product.ID)
	defer db.Close()
}
