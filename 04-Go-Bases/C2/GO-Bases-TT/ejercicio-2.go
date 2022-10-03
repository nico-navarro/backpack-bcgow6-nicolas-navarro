package main

import "fmt"

type Matrix struct {
	Values    []float64
	Height    int
	Width     int
	quadratic bool
	MaxValue  float64
}

func NewMatrix(height int, width int, values ...float64) (matrix Matrix) {
	matrix.Height = height
	matrix.Width = width
	matrix.Values = values

	matrix.quadratic = false
	if height == width {
		matrix.quadratic = true
	}

	actualMax := matrix.Values[0]
	for _, value := range values {
		if value > actualMax {
			actualMax = value
		}
	}
	matrix.MaxValue = actualMax
	return
}

func (m Matrix) Print() {
	n := 0

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			fmt.Printf("%f ", m.Values[n+j])
		}
		n = (i + 1) * m.Width
		fmt.Print("\n")
	}
}

func main() {
	matrix := NewMatrix(3, 2, 1, 2, 3, 4, 5, 6)
	matrix.Print()
}
