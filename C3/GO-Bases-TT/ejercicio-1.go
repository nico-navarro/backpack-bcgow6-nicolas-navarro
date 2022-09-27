package main

import "fmt"

type User struct {
	Name     string
	Lastname string
	Age      int
	Email    string
	Pass     string
}

func (user *User) changeName(newName string) {
	user.Name = newName
}

func (user *User) changeAge(newAge int) {
	user.Age = newAge
}

func (user *User) changeEmail(newEmail string) {
	user.Email = newEmail
}

func (user *User) changePass(newPass string) {
	user.Pass = newPass
}

func main() {
	nico := User{"Nico", "Navarro", 24, "nico.navarro@mercadolibre.cl", "secreto"}

	nico.changeName("Nicolas")
	nico.changeAge(25)
	nico.changeEmail("nicolas.navarro@mercadolibre.cl")
	nico.changePass("granSecreto")

	fmt.Printf("Nombre: %s\n", nico.Name)
	fmt.Printf("Edad: %d\n", nico.Age)
	fmt.Printf("Correo: %s\n", nico.Email)
	fmt.Printf("Pass: %s\n", nico.Pass)
}
