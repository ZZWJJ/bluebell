package routes

import (
	"bluebell/controller"
	"bluebell/gokit/endpoint"
	"bluebell/gokit/service"
	"bluebell/gokit/transport"
	"bluebell/logger"
	"net/http"

	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	kitGroup := r.Group("/kit")

	// Create a new service
	svc := service.Add{}

	// Create the endpoints
	sum := endpoint.MakeSumEndpoint(&svc)
	concat := endpoint.MakeConcatEndpoint(&svc)

	transport.MakeHttpHandler(kitGroup, sum, concat)

	// 注册
	r.POST("/signUp", controller.SignUpHandler)

	// 登录
	r.POST("/login", controller.LoginHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "错误路由",
		})
	})
	return r
}
