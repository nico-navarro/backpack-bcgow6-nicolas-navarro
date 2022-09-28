package main

import "fmt"

/*
Como puntero1 guarda el valor de la dirección de memoria de numero (&numero),
al actualizar numero y desreferenciar (*puntero1) al puntero1, puedo acceder a ese nuevo valor igual
*/
func main() {
	var puntero1 *int
	var numero int
	numero = 1
	puntero1 = &numero
	fmt.Printf("El puntero de dirección %p tiene valor %p, direccion que tiene valor de %d\n", &puntero1, puntero1, *puntero1)
	numero = 2
	fmt.Printf("El puntero de dirección %p tiene valor %p, direccion que tiene valor de %d\n", &puntero1, puntero1, *puntero1)
	fmt.Printf("El numero de dirección %p tiene valor %d, lo cual no es una dirección (por consiguiente, no puedo desreferenciarlo (*numero))\n", &numero, numero)
}
