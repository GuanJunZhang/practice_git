package controllers

import (
	// "fmt"
	"fmt"
	"getLocation/services"
	"getLocation/utils"

	// "net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func GetLocation(ctx echo.Context) error {

	validate := validator.New()
	fmt.Println("master修改get_location 第一处")
	getLocationStr := ctx.QueryParam("location")
	//参数校验
	// fmt.Println("长度=",len(getLocationStr))
	err := validate.Var(len(getLocationStr), "max=60,min=3")
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入校验失败", err.Error()))
	}

	// fmt.Println("getLocationStr=",getLocationStr)
	res := services.GetLocation(getLocationStr)
	if res == *new(utils.PCALocation) {
		return ctx.JSON(utils.Fail("Fail", res))
	} else {
		return ctx.JSON(utils.Succ("success", res))
	}
}

//可以自定义函数验证
