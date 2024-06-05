package services

import (
	"context"
	"net/http"
	"testing"

	"github.com/anv01208/online-flower-shop/auth-svc/pkg/db"
	"github.com/anv01208/online-flower-shop/auth-svc/pkg/pb"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	s := &Server{
		H: db.Init("postgres://postgres:5656@localhost/auth_svc?sslmode=disable"),
	}
	// Test case 1: Successful retrieval
	t.Run("Success", func(t *testing.T) {
		res, err := s.Register(context.Background(), &pb.RegisterRequest{
			Email:      "test@mail.com",
			Password:   "1234",
			Privileges: "seller",
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusCreated, int(res.Status))
	})

	// Test case 2: Product not found
	t.Run("SameEmail", func(t *testing.T) {
		res, err := s.Register(context.Background(), &pb.RegisterRequest{
			Email:      "test@mail.com",
			Password:   "1234",
			Privileges: "seller",
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusConflict, int(res.Status))
		assert.Equal(t, "E-Mail already exists", res.Error)
	})
}
