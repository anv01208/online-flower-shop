package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/anv01208/online-flower-shop/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.FindOne(context.Background(), &pb.FindOneProductRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
