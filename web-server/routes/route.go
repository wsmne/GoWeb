package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-server/controllers"
	"web-server/middleware"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		},
		)
	})
	r.POST("/user", controllers.Regist)
	r.GET("login")
	r.Use(middleware.ParseToken)
	r.GET("/user", controllers.GetUserByID)
	r.PUT("/user")
	r.DELETE("/user")
}
