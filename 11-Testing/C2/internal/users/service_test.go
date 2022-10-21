package users

import (
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
