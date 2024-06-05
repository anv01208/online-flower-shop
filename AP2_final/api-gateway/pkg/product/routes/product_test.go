package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/catalog/", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDkwNjgzMjIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MywiRW1haWwiOiJhbmFiZWtvdkBnbWFpbC5jb20ifQ.Sjylcsv9RkE7BvBn-mLzU58FfO6mjOyhT3ppC-ujuwQ"
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error making http request: %s\n", err)
	}
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCreateProduct(t *testing.T) {
	requestBody := CreateProductRequestBody{
		Name:     "noik",
		Stock:    1000,
		Price:    5699,
		Category: "Chinese",
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/catalog/", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTgzODQ1NTEsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6OSwiRW1haWwiOiJudXJkYUBtdXNrLmNvbSJ9.VWiQZDyno7Ims71bQ-aGGMMaaiJ5hpB6Zm557Gn9Qiw"
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error making http request: %s\n", err)
	}
	assert.Equal(t, http.StatusCreated, res.StatusCode)
}

func TestFindOneProduct(t *testing.T) {
	productId := 13

	params := httprouter.Params{{Key: "id", Value: strconv.FormatInt(int64(productId), 10)}}
	ctx := context.WithValue(context.Background(), httprouter.ParamsKey, params)

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/catalog/", nil)
	if err != nil {
		t.Fatal(err)
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTgzODQzOTgsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MTAsIkVtYWlsIjoiYmluYXJhQG11c2suY29tIn0.zAsj4bzTl54E0GsU20N0ScESex6xpfp4wrdIKbBquUI"
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error making http request: %s\n", err)
	}
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
