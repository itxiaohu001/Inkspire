package routers

import (
	"Inkspire/controllers"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	uc *controllers.UserController
}

func NewUserRouter(uc *controllers.UserController) *UserRouter {
	return &UserRouter{uc: uc}
}

func (ur *UserRouter) Init(public, private *gin.RouterGroup) {
	userPublic := public.Group("user")
	userPrivate := private.Group("user")
	{
		userPublic.POST("login", ur.uc.Login)
	}
	{
		userPrivate.GET("list", ur.uc.GetUserList)
		userPrivate.POST("register", ur.uc.Register)
	}
}
