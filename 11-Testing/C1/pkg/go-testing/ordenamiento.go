package operaciones

import "sort"

// Función que recibe un slice de enteros y los ordena
func Ordenar(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
