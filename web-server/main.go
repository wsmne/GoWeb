package main

import (
	"github.com/gin-gonic/gin"
	"web-server/logger"
	"web-server/models"
)

func main() {

	err := models.InitDataBase(logger.Log)
	if err != nil {
		logger.Log.Error("初始化数据库失败 err: " + err.Error())
	}
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
		})
	})

	router.Run(":5001")
}
