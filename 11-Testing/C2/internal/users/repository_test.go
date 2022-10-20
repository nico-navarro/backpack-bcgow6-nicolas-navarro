package users

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

func (s StubStore) Read(data interface{}) error {
	// Versión 1: usando type assertion y desreferenciando
	users := data.(*[]User)
	stubData := []User{
		{
			Id:     1,
			Name:   "Nico",
			Email:  "nico@mercadolibre.cl",
			Age:    1,
			Height: 1293,
			Active: true,
			Date:   "25-11-2012",
		},
		{
			Id:     2,
			Name:   "Gabi",
			Email:  "gabi@mercadolibre.cl",
			Age:    1,
			Height: 1293,
			Active: true,
			Date:   "25-11-2012",
		},
	}
	*users = stubData
	return nil

	// Versión 2: json Marshall y Unmarshall
	// stubData := []User{
	// 	{
	// 		Id:     1,
	// 		Name:   "Nico",
	// 		Email:  "nico@mercadolibre.cl",
	// 		Age:    1,
	// 		Height: 1293,
	// 		Active: true,
	// 		Date:   "25-11-2012",
	// 	},
	// 	{
	// 		Id:     2,
	// 		Name:   "Gabi",
	// 		Email:  "gabi@mercadolibre.cl",
	// 		Age:    1,
	// 		Height: 1293,
	// 		Active: true,
	// 		Date:   "25-11-2012",
	// 	},
	// }
	// stubDataB, _ := json.Marshal(stubData)
	// json.Unmarshal(stubDataB, data)
	// return nil
}

func (s StubStore) Write(data interface{}) error {
	stubData := []User{
		{
			Id:     1,
			Name:   "Nicolas",
			Email:  "nico@mercadolibre.cl",
			Age:    1,
			Height: 1293,
			Active: true,
			Date:   "25-11-2012",
		},
		{
			Id:     2,
			Name:   "Gabi",
			Email:  "gabi@mercadolibre.cl",
			Age:    1,
			Height: 1293,
			Active: true,
			Date:   "25-11-2012",
		},
	}
	stubDataB, _ := json.Marshal(stubData)
	json.Unmarshal(stubDataB, data)
	return nil
}

func TestGetAll(t *testing.T) {
	//arrange
	myStubStore := StubStore{}
	repo := NewRepository(myStubStore)
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
		{
			Id:     2,
			Name:   "Gabi",
			Email:  "gabi@mercadolibre.cl",
			Age:    1,
			Height: 1293,
			Active: true,
			Date:   "25-11-2012",
		},
	}
	//act
	resultado, _ := repo.GetAll()
	//assert
	assert.Equal(t, dataEsperada, resultado)
}

func TestUpdateName(t *testing.T) {
	//arrange
	myStubStore := StubStore{}
	repo := NewRepository(myStubStore)
	dataEsperada := User{1, "Nicolas", "nico@mercadolibre.cl", 1, 1293, true, "25-11-2012"}
	//act
	resultado, _ := repo.Update(1, "Nicolas", "nico@mercadolibre.cl", 1, 1293, true, "25-11-2012")
	//assert
	assert.Equal(t, dataEsperada, resultado)
}
