/*
 * @Descripttion:
 * @version:
 * @Author: fmy1993
 * @Date: 2021-04-23 08:20:51
 * @LastEditors: fmy1993
 * @LastEditTime: 2021-04-25 16:34:29
 */
package routers

import (
	"github.com/fmy1993/go-gin-admain/pkg/setting"
	// 导入了router router下的子文件夹就直接 (作为一个包)可用了
	v1 "github.com/fmy1993/go-gin-admain/routers/api/v1"
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
	// r.GET("test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"msg": "test"})
	// })
	/*
			Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
		For example, all the routes that use a common middleware for authorization could be grouped.
	*/
	apiv1 := r.Group("api/v1")

	apiv1.GET("/tags", v1.GetTags)
	apiv1.POST("/tags", v1.AddTag)
	apiv1.PUT("/tags/:id", v1.EditTag)
	apiv1.DELETE("/tags/:id", v1.DeleteTag)

	apiv1.GET("/article", v1.GetArticleList)
	apiv1.GET("/article/:id", v1.GetArticleById)
	apiv1.PUT("/article/:id", v1.EditArticle)
	apiv1.DELETE("artile/:id", v1.DeleteArticle)

	return r
}
