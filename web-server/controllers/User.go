package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"web-server/logger"
	"web-server/middleware"
	"web-server/models"
)

func GetUserByID(ctx *gin.Context) {
	//if _, exist := ctx.Get("user"); exist == false {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"code": "15006",
	//		"msg":  "无权限",
	//	})
	//	return
	//}
	id := ctx.Query("id")
	user, err := models.GetUserByID(cast.ToUint(id))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "15001",
			"msg":  "无该用户",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"user": user,
		},
	})
}

func Login(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 15004,
			"msg":  "输入错误",
		})
		return
	}
	login, err := models.GetUserByID(user.ID)
	if err != nil {
		logger.Log.Info("用户不存在，err : " + err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": 15005,
			"msg":  "用户不存在",
		})
		return
	}
	if user.UserPW == login.UserPW {
		token, _ := middleware.GenToken(user)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "登录成功",
			"data": gin.H{
				"token": token,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 15006,
			"msg":  "密码错误",
		})
		return
	}

}

func Regist(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 15003,
			"msg":  "输入错误",
		})
		return
	}
	err = models.CreateUser(user)
	if err != nil {
		logger.Log.Info("创建用户错误，err : " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 15002,
			"msg":  "创建失败",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "创建成功",
	})
}
