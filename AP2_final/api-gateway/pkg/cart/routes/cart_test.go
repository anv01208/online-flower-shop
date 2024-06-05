package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetAllCart(t *testing.T) {
	userId := 9
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/cart", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "userId", int64(userId))
	req = req.WithContext(ctx)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTgzODQzOTgsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MTAsIkVtYWlsIjoiYmluYXJhQG11c2suY29tIn0.zAsj4bzTl54E0GsU20N0ScESex6xpfp4wrdIKbBquUI"
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error making http request: %s\n", err)
	}
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestAddToCart(t *testing.T) {
	userId := 10
	requestBody := AddToCartRequestBody{
		ProductId: 1,
		Quantity:  1,
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/cart", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "userId", int64(userId))
	req = req.WithContext(ctx)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTgzODQzOTgsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MTAsIkVtYWlsIjoiYmluYXJhQG11c2suY29tIn0.zAsj4bzTl54E0GsU20N0ScESex6xpfp4wrdIKbBquUI"
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error making http request: %s\n", err)
	}
	assert.Equal(t, http.StatusCreated, res.StatusCode)
}
