package routes

import (
	"github.com/gin-gonic/gin"
	"web-server/controllers"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/user", controllers.GetUserByID)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/user")
	r.DELETE("/user")
}
