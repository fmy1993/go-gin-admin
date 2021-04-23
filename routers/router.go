package routers

import (
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())   //所谓的中间件也就是一些写好的可插拔的小的功能模块
	r.Use(gin.Recovery()) //其实就是将程序原来处理的panic错误转化为500错误直接返回给客户端
	//Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.

	// gin.Default() gin框架默认返回的engine就是带有logger 和 recover的
	gin.SetMode(setting.RunMode)
	// 在路由引擎做一些配置，然后返回
	r.GET("test", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "test"})
	})
	return r
}
