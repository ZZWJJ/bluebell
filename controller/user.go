package controller

import (
	"bluebell/model"
	"bluebell/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 用户注册接口
// @Summary 用户注册接口
// @Description 用户注册
// @Tags 用户相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query model.ParamSignUp true "查询参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /signUp [post]
func SignUpHandler(c *gin.Context) {
	// 参数校验
	p := new(model.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("sign up with wrong param ", zap.Error(err))
		// 判断是不是内置的错误类型
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"msg": "请求参数有误"})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": removeTopStruct(err.Translate(trans))})
		}
		return
	}
	// 业务处理
	if err := service.SignUp(p); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"msg": "注册失败"})
	}
	// 返回响应
	c.JSON(http.StatusOK, "ok")
}

// LoginHandler 用户登录接口
// @Summary 用户登录
// @Description 用户登录
// @Tags 用户相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query model.ParamLogin true "查询参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	// 1.获取请求参数并校验
	var p = new(model.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("login with wrong param ", zap.Error(err))
		// 判断是不是内置的错误类型
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"msg": "请求参数有误"})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": removeTopStruct(err.Translate(trans))})
		}
		return
	}
	// 2.业务逻辑处理
	token, err := service.Login(p)
	if err != nil {
		ResponseError(c, CodeInvalidPassword)
	}
	// 3.返回响应
	c.JSON(http.StatusOK, token)
}
