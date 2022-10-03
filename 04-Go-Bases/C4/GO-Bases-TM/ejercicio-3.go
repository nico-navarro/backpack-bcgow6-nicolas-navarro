package main

import (
	"fmt"
)

func checkTax(salary int) (err error) {
	if salary < 150_000 {
		err = fmt.Errorf("error: el mínimo imponible es de 150000 y el salario ingresado es de: %d", salary)
	}
	return
}

func main() {
	var salary int
	salary = 100_000
	err := checkTax(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
