package main

import (
	"Inkspire/database"
	"Inkspire/routers"
	"Inkspire/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DB_USER, utils.DB_PASS, utils.DB_HOST, utils.DB_PORT, utils.DB_DATABASE)
	database.InitDB(dsn)

	r := gin.Default()

	routers.InitEnter(r)

	if err := r.Run(":" + utils.SERVER_PORT); err != nil {
		log.Fatal(err)
	}
}
