package routers

import (
	"Inkspire/controllers"
	"Inkspire/dao"
	"Inkspire/database"
	"Inkspire/service"

	"github.com/gin-gonic/gin"
)

const v1Prefix = "/api/v1"

type Enter struct {
	UserRouters *UserRouters
}

func InitEnter(r *gin.Engine) {
	userRouters := NewUserRouters(controllers.NewUserController(
		service.NewUserService(dao.NewUserDAO(database.GetDB())),
	))

	userRouters.Init(r.Group(v1Prefix))
}
