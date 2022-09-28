package main

import "fmt"

type myCustomError struct {
}

// esto es necesario para que myCustomError cumpla con la interface default de Error
func (e *myCustomError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func checkTax(salary int) (err error) {
	if salary < 150_000 {
		err = &myCustomError{}
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
