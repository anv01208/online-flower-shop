package routes

import (
	"context"
	"net/http"

	"github.com/anv01208/online-flower-shop/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Privileges string `json:"privileges"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:      body.Email,
		Password:   body.Password,
		Privileges: body.Privileges,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
