package operaciones

import "errors"

// Errors
var (
	ErrDenominadorCero = errors.New("el denominador no puede ser 0")
)

// Función que recibe dos enteros y retorna la suma resultante
func Sumar(num1, num2 int) int {
	return num1 + num2
}

// Función que recibe dos enteros y retorna la resta o diferencia resultante
func Restar(num1, num2 int) int {
	return num1 - num2
}

// Función que recibe dos enteros y retorna la división resultante
func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, ErrDenominadorCero
	}
	return num / den, nil
}
