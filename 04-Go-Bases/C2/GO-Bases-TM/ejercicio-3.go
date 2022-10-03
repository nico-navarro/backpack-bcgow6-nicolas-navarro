package main

import (
	"fmt"
)

const (
	A = "A"
	B = "B"
	C = "C"
)

func salary(minutes int, category string) (salary float64) {
	hours := float64(minutes) / 60
	if category == A {
		salary = 3000 * hours
		salary += 0.5 * salary
	} else if category == B {
		salary = 1500 * hours
		salary += 0.2 * salary
	} else if category == C {
		salary = 1000 * hours
	}
	return
}

func main() {
	salary := salary(300, B)
	fmt.Printf("El salario es %f\n", salary)
}
