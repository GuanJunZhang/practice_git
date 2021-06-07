package utils

import (
	// "fmt"
	"getLocation/conf"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/zxysilent/logs"
)

// midAuth 登录认证中间件
func MidAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// fmt.Println("ctx.Request().URL=",ctx.Request().URL)
		// fmt.Println("ctx.RealIP()=",ctx.RealIP())
		// fmt.Println("ctx.Path()=",ctx.Path())
		// fmt.Println("conf.App.Jwt.LoginPath",conf.App.Jwt.LoginPath)
		if ctx.Request().URL.Path == conf.App.Jwt.LoginPath {
			return next(ctx)
		}
		tokenRaw := ctx.Request().Header.Get("token")
		logs.Info("tokenRaw=", tokenRaw)
		if tokenRaw == "" {
			return ctx.JSON(ErrJwt("token不可为空"))
		}
		claims, err := parseAuthToken(tokenRaw)
		logs.Info("claims=", claims, "err=", err)
		if err != nil {
			return ctx.JSON(ErrJwt("请重新登陆", err.Error()))
		}
		if ctx.RealIP() != claims.Ip {
			return ctx.JSON(ErrJwt("网络变更,请重新登陆"))
		}
		ctx.Set("uid", claims.Uid) //存到了store字段里面
		// ctx.Set("myUid",1234)//
		logs.Info("ctx=", ctx) //
		// fmt.Println("next(ctx)=",next(ctx))
		return next(ctx)
	}
}

type authClaims struct {
	Uid int
	Ip  string
	jwt.StandardClaims
}

func CreateAuthToken(uid int, ip string) (string, error) {
	nowSecond := int64(time.Now().Unix())
	// fmt.Println("nowSecond=",nowSecond)
	expireAtSecond := nowSecond + int64(conf.App.Jwt.AuthLifetime) //加了两小时,conf.App.Jwt.AuthLifetime= 7200
	// fmt.Println("conf.App.Jwt.AuthLifetime=",conf.App.Jwt.AuthLifetime)
	// fmt.Println("expireAtSecond=",expireAtSecond)
	claims := &authClaims{
		Uid: uid,
		Ip:  ip,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAtSecond,
			NotBefore: nowSecond,
		},
	}
	// fmt.Println("conf.App.Jwt.AuthKey=",conf.App.Jwt.AuthKey)//conf.App.Jwt.AuthKey= youshangjiao
	return CreateToken(claims, conf.App.Jwt.AuthKey)
}

func parseAuthToken(tokenStr string) (*authClaims, error) {
	claims, err := ParseToken(tokenStr, &authClaims{}, conf.App.Jwt.AuthKey)
	if err != nil {
		return nil, err
	}
	if claims, ok := claims.(*authClaims); ok {
		// fmt.Println("claims.(*authClaims).Uid",claims.Uid)
		// fmt.Println("claims.(*authClaims).Ip",claims.Ip)
		// fmt.Println("claims.StandardClaims.Audience",claims.StandardClaims.Audience)
		// fmt.Println("claims.StandardClaims.ExpiresAt",claims.StandardClaims.ExpiresAt)
		// fmt.Println("claims.StandardClaims.Id",claims.StandardClaims.Id)
		// fmt.Println("claims.StandardClaims.IssuedAt",claims.StandardClaims.IssuedAt)
		// fmt.Println("claims.StandardClaims.Issuer",claims.StandardClaims.Issuer)
		// fmt.Println("claims.StandardClaims.NotBefore",claims.StandardClaims.NotBefore)
		// fmt.Println("claims.StandardClaims.Subject",claims.StandardClaims.Subject)
		return claims, nil
	}

	return nil, err //理论上一定执行不到
}
