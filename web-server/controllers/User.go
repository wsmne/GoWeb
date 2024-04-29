package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-server/models"
)

func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user := models.GetUserByID(id)
	ctx.JSON(200, gin.H{
		"code": 0,
		"user": user,
	})
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"user": user,
		})
		return
	}
	err = models.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"user": user,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
	})
}
