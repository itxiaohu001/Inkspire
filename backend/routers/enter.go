package routers

import (
	"Inkspire/controllers"
	"Inkspire/dao"
	"Inkspire/database"
	"Inkspire/middlewares"
	"Inkspire/service"

	"github.com/gin-gonic/gin"
)

const v1Prefix = "/api/v1"

type Enter struct {
	UserRouters *UserRouter
}

func InitEnter(r *gin.Engine) {
	root := r.Group(v1Prefix)

	private := root
	public := root

	private.Use(middlewares.AuthMiddleware())

	userRouter := NewUserRouter(controllers.NewUserController(
		service.NewUserService(dao.NewUserDAO(database.GetDB())),
	))

	userRouter.Init(public, private)
}
