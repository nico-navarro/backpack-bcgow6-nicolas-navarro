package main

import (
	"errors"
	"fmt"
	"os"
)

func calculateBonus(bestSalary float64, workedMonths int) (bonus float64, err error) {
	if bestSalary < 0 || workedMonths < 0 {
		err = errors.New("error: no puedes ingresar números negativos")
		return
	}
	bonus = (bestSalary / 12) * float64(workedMonths)
	return
}

func calculateSalary(hours float64, hourPay float64) (salary float64, err error) {
	salary = hours * hourPay
	if salary >= 150_000 {
		salary -= salary * 0.1
	}
	if hours < 80 {
		err = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	return
}

func main() {
	var hours float64 = 90
	var hourPay float64 = 60
	salary, err := calculateSalary(hours, hourPay)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bonus, err := calculateBonus(salary, 6)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("El medio aguinaldo es: %f\n", bonus)
}
