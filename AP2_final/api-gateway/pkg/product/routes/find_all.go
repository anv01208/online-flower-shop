package routes

import (
	"context"
	"io"
	"net/http"

	pb "github.com/anv01208/online-flower-shop/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func FindAll(ctx *gin.Context, c pb.ProductServiceClient) {
	stream, err := c.FindAll(context.Background(), &pb.Empty{})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	var products []*pb.ProductData

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return
		}

		product := res.GetData()
		products = append(products, product...)
	}

	// Return the products as a JSON response
	ctx.JSON(http.StatusOK, products)
}
