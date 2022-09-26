package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunction(values ...int) (minValue float64) {
	actualMin := values[0]
	for _, value := range values {
		if value < actualMin {
			actualMin = value
		}
	}
	minValue = float64(actualMin)
	return
}

func avgFunction(values ...int) (avgValue float64) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	totalValues := len(values)
	avgValue = float64(sum) / float64(totalValues)
	return
}

func maxFunction(values ...int) (maxValue float64) {
	actualMax := values[0]
	for _, value := range values {
		if value > actualMax {
			actualMax = value
		}
	}
	maxValue = float64(actualMax)
	return
}

func operation(operation string) (func(...int) float64, error) {
	switch operation {
	case minimum:
		return minFunction, nil
	case average:
		return avgFunction, nil
	case maximum:
		return maxFunction, nil
	}
	err := errors.New("No está definida esa operación")
	return nil, err
}

func main() {
	// minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	// maxFunc, err := operation(maximum)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		// minValue := minFunc(10, 3, 3, 4, 10, 2, 1, 5)
		averageValue := averageFunc(2, 10, 3, 4, 1, 2, 4, 5)
		// maxValue := maxFunc(2, 3, 15, 4, 1, 2, 4, 5)
		fmt.Printf("El valor de la operación es: %f\n", averageValue)
	}
}
