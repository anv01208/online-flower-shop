package routes

import (
	"context"
	"net/http"

	"github.com/anv01208/online-flower-shop/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Privileges string `json:"privileges"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:      b.Email,
		Password:   b.Password,
		Privileges: b.Privileges,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
