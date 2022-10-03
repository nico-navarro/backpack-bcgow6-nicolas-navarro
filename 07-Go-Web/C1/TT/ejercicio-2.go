// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
// Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

package main

import "fmt"

func main() {
	var age int = 15
	var hired bool = true
	var antiquity int = 3
	var salary int = 145000
	if age > 22 && hired && antiquity > 1 {
		fmt.Println("Tienes acceso al crédito")
		if salary > 100000 {
			fmt.Println("No pagarás intereses")
		} else {
			fmt.Println("Pagarás intereses")
		}
	} else {
		fmt.Println("No tienes acceso al crédito")
	}
}
