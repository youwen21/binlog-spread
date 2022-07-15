package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AdminToken 如果是已登录用户， 从token中提取用户ID
func AdminToken(c *gin.Context) {
	err := AdminCheckToken(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		c.Abort()
	}
	c.Next()
}

//AdminCheckToken 如果是已登录用户， 从token中提取用户ID
// 提取这个方法是为NoRoute中REST使用
func AdminCheckToken(c *gin.Context) error {
	tokenString := c.GetHeader("AdminAuthorization")
	if tokenString == "" {
		tokenString, _ = c.Cookie("AdminAuthorization")
	}
	if tokenString == "" {
		return errors.New("缺失认证token")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := os.Getenv("ADMIN_JWT_SECRET")
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSampleSecret), nil
	})

	if nil != err {
		return errors.New("admin-token解析token失败")
	}

	if !token.Valid {
		return errors.New("admin-token解析无效")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("admin-token, claims无效")
	}
	c.Set("admin_id", claims["uid"])

	return nil
}
