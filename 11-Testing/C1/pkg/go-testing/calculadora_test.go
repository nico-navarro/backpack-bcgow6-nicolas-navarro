package operaciones

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	// Arrange: Se inicializan los datos a usar en el test (input/output)
	num1 := 10
	num2 := 3
	resultadoEsperado := 7
	// Act: Se ejecuta el test
	resultado := Restar(num1, num2)
	// Assert: Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestDividir(t *testing.T) {
	// Arrange: Se inicializan los datos a usar en el test (input/output)
	num1 := 10
	num2 := 0
	resultadoEsperado := ErrDenominadorCero
	// Act: Se ejecuta el test
	_, err := Dividir(num1, num2)
	// Assert: Se validan los resultados
	assert.ErrorIs(t, err, resultadoEsperado)
}
