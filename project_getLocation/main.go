package main

import (
	"context"
	"fmt"
	"getLocation/conf"
	"getLocation/route"
	"getLocation/utils"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/zxysilent/logs"
)

func main() {
	fmt.Println("master修改main 第一处")
	fmt.Println("master修改main 第二处")
	fmt.Println("master修改main 第三处")
	fmt.Println("master修改main 第四处")
	conf.Init()
	fmt.Println("zgj_test修改main")
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(utils.MidAuth)
	fmt.Println("新的master分支修改")
	route.InitRouter(e.Router())

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	// go e.Start(":8080")
	// fmt.Println("conf.App.Http.Address=",conf.App.Http.Address)
	//改进根据配置文件配置Ip,端口，开启协程
	go func() {
		if err := e.Start(conf.App.Http.Address); err != nil {
			logs.Error(err.Error())
		}
	}()
	// 开启系统信号接收通道
	// 防止系统推出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if e != nil {
			_ = e.Shutdown(ctx)
		}
	case syscall.SIGHUP:
	default:
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
