package product

import (
	"fmt"

	"github.com/anv01208/online-flower-shop/api-gateway/pkg/config"
	"github.com/anv01208/online-flower-shop/api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}
