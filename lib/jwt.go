package lib

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenToken(uid int) (string, error) {
	hmacSampleSecret := os.Getenv("JWT_SECRET")

	// 生成token https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"uid": uid, "exp": time.Now().Unix() + 86400*30})

	return jwtToken.SignedString([]byte(hmacSampleSecret))
}

func VerifyGetToken(s string) (*jwt.Token, error) {
	token, err := jwt.Parse(s, keyFuncX)

	if nil != err {
		return nil, errors.New("解析token失败:" + s)
	}

	if !token.Valid {
		return nil, errors.New("token无效:" + s)
	}

	return token, nil
}

func keyFuncX(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	hmacSampleSecret := os.Getenv("JWT_SECRET")
	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(hmacSampleSecret), nil
}
