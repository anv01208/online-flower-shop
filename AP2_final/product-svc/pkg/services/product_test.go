package services

import (
	"context"
	"net/http"
	"testing"

	"github.com/anv01208/online-flower-shop/product-svc/pkg/db"
	"github.com/anv01208/online-flower-shop/product-svc/pkg/pb"
	"github.com/stretchr/testify/assert"
)

func TestFindOne(t *testing.T) {
	s := &Server{
		H: db.Init("postgres://postgres:5656@localhost/product_svc?sslmode=disable"),
	}
	// Test case 1: Successful retrieval
	t.Run("Success", func(t *testing.T) {
		res, err := s.FindOne(context.Background(), &pb.FindOneProductRequest{
			Id: 12,
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusOK, int(res.Status))
		assert.Equal(t, 12, int(res.Data.Id))
		assert.Equal(t, "Abibas", res.Data.Name)
	})

	// Test case 2: Product not found
	t.Run("NotFound", func(t *testing.T) {
		res, err := s.FindOne(context.Background(), &pb.FindOneProductRequest{
			Id: 134,
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusNotFound, int(res.Status))
		assert.Equal(t, "record not found", res.Error)
	})
}

func TestCreateProduct(t *testing.T) {
	s := &Server{
		H: db.Init("postgres://postgres:5656@localhost/product_svc?sslmode=disable"),
	}

	// Test case 1: Successful creation
	t.Run("Success", func(t *testing.T) {
		req := &pb.CreateProductRequest{
			Name:     "Test Product",
			Stock:    10,
			Price:    10,
			Category: "Test Category",
		}

		res, err := s.CreateProduct(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusCreated, int(res.Status))
		assert.NotEmpty(t, res.Id)
	})
}
