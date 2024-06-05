package product

import (
	"github.com/anv01208/online-flower-shop/api-gateway/pkg/auth"
	"github.com/anv01208/online-flower-shop/api-gateway/pkg/config"
	"github.com/anv01208/online-flower-shop/api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/catalog")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/", svc.FindAll)
	routes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) FindAll(ctx *gin.Context) {
	routes.FindAll(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
