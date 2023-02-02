package controllors

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SingUpHandler 处理请求的函数
func SingUpHandler(c *gin.Context) {
	//1、参数校验
	p := new(models.ParamSignUp)
	//只能判断参数类型是不是正确的
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误 直接返回响应
		zap.L().Error("SingUp with invalid param", zap.Error(err))

		//判断err是 validator.validatorErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//详细的判断
	//手动的对请求参数进行详细的业务规则校验
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	// 	zap.L().Error("SingUp with invalid param")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请求参数有误",
	// 	})
	// 	return
	// }
	fmt.Println(p)
	//2、业务处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3、返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	//1、获取请求参数以及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误 直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))

		//判断err是 validator.validatorErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2、业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3、返回响应
	ResponseSuccess(c, token)
}
