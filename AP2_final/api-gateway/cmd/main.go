package main

import (
	"log"

	"github.com/anv01208/online-flower-shop/api-gateway/pkg/auth"
	"github.com/anv01208/online-flower-shop/api-gateway/pkg/cart"
	"github.com/anv01208/online-flower-shop/api-gateway/pkg/config"
	"github.com/anv01208/online-flower-shop/api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	cart.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
