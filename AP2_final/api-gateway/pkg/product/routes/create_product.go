package routes

import (
	"context"
	"net/http"

	"github.com/anv01208/online-flower-shop/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

type CreateProductRequestBody struct {
	Name     string `json:"name"`
	Stock    int64  `json:"stock"`
	Price    int64  `json:"price"`
	Category string `json:"category"`
}

type CustomError struct {
	StatusCode int
	Message    string
}

func (e CustomError) Error() string {
	return e.Message
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	priv, _ := ctx.Get("priv")

	if priv.(string) != "seller" {
		ctx.AbortWithError(http.StatusForbidden, CustomError{
			StatusCode: http.StatusForbidden,
			Message:    "Access forbidden",
		})
		return
	}

	body := CreateProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:     body.Name,
		Stock:    body.Stock,
		Price:    body.Price,
		Category: body.Category,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
