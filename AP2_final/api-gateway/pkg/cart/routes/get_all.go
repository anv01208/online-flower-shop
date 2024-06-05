package routes

import (
	"context"
	"io"
	"net/http"

	"github.com/anv01208/online-flower-shop/api-gateway/pkg/cart/pb"
	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context, c pb.CartServiceClient) {
	userId, _ := ctx.Get("userId")
	//fmt.Println("userId =", userId)

	stream, err := c.GetCartItems(context.Background(), &pb.ViewCartRequest{
		UserId: userId.(int64),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	var cartItems []*pb.CartItem

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return
		}

		product := res.GetCartItem()
		cartItems = append(cartItems, product...)
	}

	// Return the products as a JSON response
	ctx.JSON(http.StatusOK, cartItems)
}
