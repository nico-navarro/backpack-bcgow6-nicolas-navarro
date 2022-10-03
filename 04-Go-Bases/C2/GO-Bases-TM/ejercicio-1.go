package main

import "fmt"

func calculateTax(salary int) float64 {
	var tax float64 = 0
	if salary > 50000 {
		tax = 0.17
		if salary > 150000 {
			tax += 0.1
		}
	}
	return tax
}

func main() {
	var salary int = 60000
	var tax float64 = calculateTax(salary)
	fmt.Printf("El impuesto correspondiente es %f\n", tax)
}
