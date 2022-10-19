package operaciones

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	// Arrange: Se inicializan los datos a usar en el test (input/output)
	numbers := []int{7, 5, 10, 1}
	resultadoEsperado := []int{1, 5, 7, 10}
	// Act: Se ejecuta el test
	resultado := Ordenar(numbers)
	// Assert: Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
