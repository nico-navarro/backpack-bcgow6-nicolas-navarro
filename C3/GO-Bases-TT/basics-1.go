package main

import "fmt"

/*
Esto da error porque no hay nada que desferenciar aun,
ya que puntero1 nunca recibió un valor que pudiese desferenciar después.
*/

func main() {
	var puntero1 *int
	*puntero1 = 10
	fmt.Printf("El puntero de dirección %p tiene valor %p, direccion que tiene valor de %d\n", &puntero1, puntero1, *puntero1)
}
