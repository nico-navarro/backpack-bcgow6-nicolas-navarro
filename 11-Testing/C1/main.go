package main

// Se importa el package testing
import (
	calculadora "C1/pkg/go-testing"
	"fmt"
)

func main() {
	suma := calculadora.Sumar(2, 3)
	fmt.Println(suma)
}
