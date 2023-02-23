package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册
	r.POST("signUp", controller.SignUpHandler)

	// 登录
	r.POST("login", controller.LoginHandler)

	r.GET("ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "错误路由",
		})
	})
	return r
}
