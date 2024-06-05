package routes

import (
	"context"
	"net/http"

	"github.com/anv01208/online-flower-shop/api-gateway/pkg/cart/pb"
	"github.com/gin-gonic/gin"
)

type AddToCartRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

func AddToCart(ctx *gin.Context, c pb.CartServiceClient) {
	b := AddToCartRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId, _ := ctx.Get("userId")

	res, err := c.AddToCart(context.Background(), &pb.AddToCartRequest{
		ProductId: b.ProductId,
		Quantity:  int64(b.Quantity),
		UserId:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
