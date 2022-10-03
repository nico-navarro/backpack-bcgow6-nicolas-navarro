// Según el siguiente mapa, ayuda a imprimir la edad de Benjamin.
// Por otro lado también es necesario:
// - Saber cuántos de sus empleados son mayores de 21 años.
// - Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// - Eliminar a Pedro del mapa.

package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// Edad de Benjamin
	ageBenjamin := employees["Benjamin"]
	fmt.Printf("La edad de Benjamín es %d\n", ageBenjamin)

	// Empleados mayores a 21
	var seniors int = 0
	for _, age := range employees {
		if age > 21 {
			seniors++
		}
	}
	fmt.Printf("La cantidad de empleados mayores a 21 es: %d\n", seniors)

	//Agregar un empleado a la lista
	employees["Federico"] = 25

	// Eliminar a Pedro del map
	delete(employees, "Pedro")
}
