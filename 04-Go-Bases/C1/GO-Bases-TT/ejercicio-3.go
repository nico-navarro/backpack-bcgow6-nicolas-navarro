// Realizar una aplicación que contenga una variable con el número del mes.
// 1. Según el número, imprimir el mes que corresponda en texto.
// 2. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por
// qué?
// Ej: 7, Julio

package main

import "fmt"

func main() {
	var months = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
	var selection = 5
	fmt.Printf("El mes %d es %s\n", selection, months[selection])
}
