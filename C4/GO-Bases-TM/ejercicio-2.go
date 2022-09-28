package main

import (
	"errors"
	"fmt"
)

func checkTax(salary int) (err error) {
	if salary < 150_000 {
		err = errors.New("error: el salario ingresado no alcanza el mínimo imponible")
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
