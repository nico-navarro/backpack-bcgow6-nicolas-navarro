package main

import (
	"errors"
	"fmt"
)

func addPositiveNumbers(numbers []int) (int, error) {
	var total int = 0
	for _, number := range numbers {
		if number < 0 {
			return 0, errors.New("Ingresaste un número negativo")
		}
		total += number
	}
	return total, nil
}

func calculateMean(scores ...int) (float64, error) {
	sum, err := addPositiveNumbers(scores)
	if err != nil {
		return 0, err
	} else {
		var mean float64
		total := len(scores)
		mean = float64(sum) / float64(total)
		return mean, nil
	}
}

func main() {
	mean, err := calculateMean(7, 6, 1, 3, 4, 5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("El promedio es %f\n", mean)
	}
}
