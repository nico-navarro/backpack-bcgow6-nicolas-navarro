package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubRepo struct {
	stubData interface{}
}

func (s StubRepo) GetAllBySeller(sellerId string) ([]Product, error) {
	return []Product{}, nil
}

func TestGetAllBySeller(t *testing.T) {
	//arrange
	dataEsperada := []Product{
		{
			ID:          "1",
			SellerID:    "1",
			Description: "a",
			Price:       12,
		},
	}
	myStubRepo := StubRepo{stubData: dataEsperada}
	service := NewService(myStubRepo)
	//act
	resultado, _ := service.GetAllBySeller("1")
	//assert
	assert.Equal(t, dataEsperada, resultado)
}
