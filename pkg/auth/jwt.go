package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtPayload struct {
	User interface{} `json:"user"`
	jwt.StandardClaims
}

// GenerateJwtToken 生成jwt token
func GenerateJwtToken(secret string, expire int64, user interface{}, issuer string) (string, error) {
	data := JwtPayload{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expire,
			Issuer:    issuer,
		},
	}
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err := j.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("jwt 生成token失败" + err.Error())
	}
	return token, nil
}

// ParseJwtToken 解析 jwt token
func ParseJwtToken(jwtToken string, secret string) (*JwtPayload, error) {
	if jwtToken == "" {
		return nil, errors.New("token 为空")
	}
	token, err := jwt.ParseWithClaims(jwtToken, &JwtPayload{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("jwt token 解析失败" + err.Error())
	}
	if claims, ok := token.Claims.(*JwtPayload); ok && token.Valid {
		return claims, nil
	} else {
		return claims, errors.New("jwt 解析验证后失败")
	}
}
