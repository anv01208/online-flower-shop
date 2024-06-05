package main

import (
	"fmt"
	"log"
	"net"

	"github.com/anv01208/online-flower-shop/cart-svc/pkg/client"
	"github.com/anv01208/online-flower-shop/cart-svc/pkg/config"
	"github.com/anv01208/online-flower-shop/cart-svc/pkg/db"
	"github.com/anv01208/online-flower-shop/cart-svc/pkg/pb"
	"github.com/anv01208/online-flower-shop/cart-svc/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Cart Svc on", c.Port)

	s := service.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCartServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
