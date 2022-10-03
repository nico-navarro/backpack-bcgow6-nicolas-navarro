package main

import "fmt"

type Student struct {
	Name     string
	Lastname string
	DNI      string
	Date     string
	// Colors   []string -> []string{"Azul", "Verde"}
	// food     string // private
}

// esto es solo una función que recibe el struct, NO es un método
func details(student Student) {
	// Forma rápida de imprimir un struct
	// fmt.Printf("Detalles: %+v\n", student)
	fmt.Printf("%s %s\n", student.Name, student.Lastname)
	fmt.Printf("- DNI: %s\n", student.DNI)
	fmt.Printf("- Fecha: %s\n", student.Date)
}

// método
func (student Student) detailsMethod() {
	fmt.Printf("%s %s\n", student.Name, student.Lastname)
	fmt.Printf("- DNI: %s\n", student.DNI)
	fmt.Printf("- Fecha: %s\n", student.Date)
}

func main() {
	nico := Student{
		"Nicolas", "Navarro", "19231321-k", "21-12-1997",
	}
	//función que recibe un struct Student
	details(nico)

	//metodo
	nico.detailsMethod()

	// De esta forma puedo inicializar sin pasar todos los args
	// gabi := Student{
	// 	Name: "Gabriela", Lastname: "Vergara",
	// }
	// details(gabi)
}
