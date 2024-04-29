package main

import (
	"github.com/gin-gonic/gin"
	"web-server/logger"
	"web-server/models"
	"web-server/routes"
)

func main() {

	err := models.InitDataBase(logger.Log)
	if err != nil {
		logger.Log.Error("初始化数据库失败 err: " + err.Error())
	}
	router := gin.Default()

	routes.RegisterRoute(router)

	router.Run(":5001")
}
