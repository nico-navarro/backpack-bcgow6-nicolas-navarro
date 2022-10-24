package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubRepo struct {
	stubData  []Product
	stubError error
}

func (s *StubRepo) GetAllBySeller(sellerId string) ([]Product, error) {
	if s.stubError != nil {
		return []Product{}, s.stubError
	}
	return s.stubData, nil
}

func TestGetAllBySeller(t *testing.T) {
	//arrange
	dataEsperada := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	myStubRepo := StubRepo{stubData: dataEsperada}
	service := NewService(&myStubRepo)
	//act
	resultado, _ := service.GetAllBySeller("1")
	//assert
	assert.Equal(t, dataEsperada, resultado)
}

func TestGetAllBySellerFail(t *testing.T) {
	//arrange
	myStubRepo := StubRepo{stubError: errors.New("error in repository")}
	service := NewService(&myStubRepo)
	//act
	_, err := service.GetAllBySeller("1")
	//assert
	assert.NotNil(t, err)
}
