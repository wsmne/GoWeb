package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"time"
	"web-server/models"
)

var CustomSecret = []byte("wsm的电子商城")

type CustomClaims struct {
	UserName string // 自定义字段
	Id       uint
	Type     string
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 24

func GenToken(user models.User) (string, error) {
	// 创建一个我们自己的声明
	claims := &CustomClaims{
		user.UserName, // 自定义字段
		user.ID,
		user.Type,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "wsm", // 签发人
			Subject:   "user token",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenStr, err := token.SignedString(CustomSecret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解析JWT
func ParseToken(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		ctx.Set("auth", 0)
		return
	}
	// 对token对象中的Claim进行类型断言
	Claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return // 校验token
	}
	user := &models.User{
		Model:    gorm.Model{ID: Claims.Id},
		UserName: Claims.UserName,
		Type:     Claims.Type,
	}
	ctx.Set("user", user)
	return
}
