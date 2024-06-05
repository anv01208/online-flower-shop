package service

import (
	"context"
	"net/http"
	"testing"

	"github.com/anv01208/online-flower-shop/cart-svc/pkg/db"
	"github.com/anv01208/online-flower-shop/cart-svc/pkg/pb"
	"github.com/stretchr/testify/assert"
)

func TestAddToCart(t *testing.T) {
	s := &Server{
		H: db.Init("postgres://postgres:5656@localhost/cart_svc?sslmode=disable"),
	}
	// Test case 1: Successful creation
	t.Run("Success", func(t *testing.T) {
		req := &pb.AddToCartRequest{
			ProductId: 1,
			Quantity:  100,
			UserId:    5,
		}

		res, err := s.AddToCart(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusOK, int(res.Status))
	})
}

func TestRemoveFromCart(t *testing.T) {
	s := &Server{
		H: db.Init("postgres://postgres:5656@localhost/cart_svc?sslmode=disable"),
	}
	// Test case 1: Successful retrieval
	t.Run("Success", func(t *testing.T) {
		res, err := s.RemoveFromCart(context.Background(), &pb.RemoveFromCartRequest{
			CartItemId: 1,
			UserId:     5,
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusOK, int(res.Status))
		assert.Equal(t, "", res.Error)
	})

	// Test case 2: Product not found
	t.Run("NotFound", func(t *testing.T) {
		res, err := s.RemoveFromCart(context.Background(), &pb.RemoveFromCartRequest{
			CartItemId: 1,
			UserId:     5,
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusOK, int(res.Status))
		assert.Equal(t, "", res.Error)
	})
}
