package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/anv01208/online-flower-shop/api-gateway/pkg/cart/pb"
	"github.com/gin-gonic/gin"
)

func RemoveFromCart(ctx *gin.Context, c pb.CartServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	userId, _ := ctx.Get("userId")

	res, err := c.RemoveFromCart(context.Background(), &pb.RemoveFromCartRequest{
		CartItemId: id,
		UserId:     userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
