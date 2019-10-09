/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-09 15:24:41
 * @LastEditTime: 2019-10-09 18:11:24
 * @LastEditors: Please set LastEditors
 */
package jkt

import (
	"errors"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/dgrijalva/jwt-go"
	jkt "github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	jkt.StandardClaims
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)

func NewToken(user, pwd string) string {
	claims := &JWTClaims{
		UserID:   1,
		Username: user,
		Password: pwd,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signesToken, err := getToken(claims)
	if err != nil {
		logs.Informational("token", err.Error())
		return ""
	}
	return signesToken
}

func getToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}
	return signedToken, nil
}

func VerifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	return claims, nil
}
