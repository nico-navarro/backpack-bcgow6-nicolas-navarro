package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type User struct {
	Name     string
	Lastname string
	Email    string
	Products []Product
}

func newProduct(name string, price float64) (product Product) {
	product.Name = name
	product.Price = price
	return
}

func addProduct(user *User, product Product, quantity int) {
	product.Quantity = quantity
	user.Products = append(user.Products, product)
}

func deleteProducts(user *User) {
	user.Products = []Product{}
}

func main() {
	nico := User{"Nico", "Navarro", "nicolas.navarro@mercadolibre.cl", []Product{}}
	shoe := newProduct("Nike", 250)
	tv := newProduct("TV Sony", 1550)
	addProduct(&nico, shoe, 2)
	addProduct(&nico, tv, 1)
	fmt.Printf("%+v\n", nico)
	deleteProducts(&nico)
	fmt.Printf("%+v\n", nico)
}
