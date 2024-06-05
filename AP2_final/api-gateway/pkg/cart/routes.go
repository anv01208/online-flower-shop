package cart

import (
	"github.com/anv01208/online-flower-shop/api-gateway/pkg/auth"
	"github.com/anv01208/online-flower-shop/api-gateway/pkg/cart/routes"
	"github.com/anv01208/online-flower-shop/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/cart")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.PutInTheCart)
	routes.GET("/", svc.GetAll)
	routes.DELETE("/:id", svc.RemoveFromCart)
}

func (svc *ServiceClient) PutInTheCart(ctx *gin.Context) {
	routes.AddToCart(ctx, svc.Client)
}

func (svc *ServiceClient) GetAll(ctx *gin.Context) {
	routes.GetAll(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveFromCart(ctx *gin.Context) {
	routes.RemoveFromCart(ctx, svc.Client)
}
