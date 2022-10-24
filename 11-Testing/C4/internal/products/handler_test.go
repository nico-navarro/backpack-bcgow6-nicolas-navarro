package products

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockGin(sellerID string) (*gin.Context, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	req := &http.Request{
		URL: &url.URL{},
	}
	query := req.URL.Query()
	query.Add("seller_id", sellerID)
	req.URL.RawQuery = query.Encode()
	ctx.Request = req
	return ctx, recorder
}

type StubService struct {
	stubData  []Product
	stubError error
}

func (s *StubService) GetAllBySeller(sellerID string) ([]Product, error) {
	if s.stubError != nil {
		return []Product{}, s.stubError
	}
	return s.stubData, nil
}

func TestHandlerGetAllBySeller(t *testing.T) {
	// arrange
	dataEsperada := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	statusEsperado := http.StatusOK
	myStubService := StubService{stubData: dataEsperada}
	handler := NewHandler(&myStubService)

	ctx, recorder := mockGin("1")

	// act
	handler.GetProducts(ctx)
	// parse response body
	response := recorder.Result()
	bytesBody, _ := io.ReadAll(response.Body)
	var body []Product
	json.Unmarshal(bytesBody, &body)

	// assert
	assert.Equal(t, statusEsperado, response.StatusCode)
	assert.Equal(t, dataEsperada, body)
}

func TestHandlerGetAllBySellerFail400(t *testing.T) {
	// arrange
	statusEsperado := http.StatusBadRequest
	myStubService := StubService{stubError: errors.New("internal server error")}
	handler := NewHandler(&myStubService)

	ctx, recorder := mockGin("")

	// act
	handler.GetProducts(ctx)
	// parse response body
	response := recorder.Result()
	bytesBody, _ := io.ReadAll(response.Body)
	var body []Product
	json.Unmarshal(bytesBody, &body)

	// assert
	assert.Equal(t, statusEsperado, response.StatusCode)
}

func TestHandlerGetAllBySellerFail500(t *testing.T) {
	// arrange
	statusEsperado := http.StatusInternalServerError
	myStubService := StubService{stubError: errors.New("internal server error")}
	handler := NewHandler(&myStubService)

	ctx, recorder := mockGin("1")

	// act
	handler.GetProducts(ctx)
	// parse response body
	response := recorder.Result()
	bytesBody, _ := io.ReadAll(response.Body)
	var body []Product
	json.Unmarshal(bytesBody, &body)

	// assert
	assert.Equal(t, statusEsperado, response.StatusCode)
}
