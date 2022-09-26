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

func details(student Student) {
	// Forma rápida de imprimir un struct
	// fmt.Printf("Detalles: %+v\n", student)
	fmt.Printf("%s %s\n", student.Name, student.Lastname)
	fmt.Printf("- DNI: %s\n", student.DNI)
	fmt.Printf("- Fecha: %s\n", student.Date)
}

func main() {
	nico := Student{
		"Nicolas", "Navarro", "19231321-k", "21-12-1997",
	}
	details(nico)

	// De esta forma puedo inicializar sin pasar todos los args
	// gabi := Student{
	// 	Name: "Gabriela", Lastname: "Vergara",
	// }
	// details(gabi)
}
