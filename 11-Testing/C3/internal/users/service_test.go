package users

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceUpdate(t *testing.T) {
	//arrange
	myStubStore := StubStore{}
	repo := NewRepository(&myStubStore)
	service := NewService(repo)
	dataEsperada := User{1, "Nicolas", "nico@mercadolibre.cl", 1, 1293, true, "25-11-2012"}
	//act
	resultado, _ := service.Update(1, "Nicolas", "nico@mercadolibre.cl", 1, 1293, true, "25-11-2012")
	//assert
	assert.Equal(t, dataEsperada, resultado)
	assert.True(t, myStubStore.ReadWasCalled)
}

func TestServiceDelete(t *testing.T) {

	myStubStore := StubStore{}
	repo := NewRepository(&myStubStore)
	service := NewService(repo)

	// 1. Valida que user borrado efectivamente no exista
	//	se debe modificar el Read() para que funcione bien el GetAll
	// arrange
	dataEsperada := []User{
		{
			Id:     1,
			Name:   "Nico",
			Email:  "nico@mercadolibre.cl",
			Age:    1,
			Height: 1293,
			Active: true,
			Date:   "25-11-2012",
		},
	}
	//act
	_, _ = service.Delete(1)
	users, _ := service.GetAll()
	//assert
	assert.Equal(t, dataEsperada, users)

	// 2. Obtiene el error correspondiente
	//arrange
	expectedError := errors.New("no existe el ID especificado")
	//act
	_, err := service.Delete(15)
	//assert
	assert.EqualError(t, err, expectedError.Error())
}
