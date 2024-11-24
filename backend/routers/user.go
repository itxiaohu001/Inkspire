package routers

import (
	"Inkspire/controllers"
	"Inkspire/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRouters struct {
	uc *controllers.UserController
}

func NewUserRouters(uc *controllers.UserController) *UserRouters {
	return &UserRouters{uc: uc}
}

func (ur *UserRouters) Init(router *gin.RouterGroup) {
	userGroup := router.Group("user")
	{
		userGroup.POST("login", ur.uc.Login)
	}
	{
		needToken := userGroup
		needToken.Use(middlewares.AuthMiddleware())
		needToken.GET("list", ur.uc.GetUserList)
		needToken.POST("register", ur.uc.Register)
	}
}
