package v1

import (
	"net/http"

	"github.com/fmy1993/go-gin-admain/models"
	e "github.com/fmy1993/go-gin-admain/pkg/error"
	"github.com/fmy1993/go-gin-admain/pkg/setting"
	"github.com/fmy1993/go-gin-admain/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

/*
tag.go will implement CRUD about tags
*/

func GetTags(c *gin.Context) {
	// c.JSON(200, gin.H{"msg": "send by router"})
	name := c.Query("name")
	//go的map自带动态添加属性，而且直接赋值就行，就像无限长的hashtable
	//maps:sql的查询条件
	maps := make(map[string]interface{})
	//data:查询数据的集合
	data := make(map[string]interface{})
	// check http parameter is not null ,other check can be done in front end
	if name != "" {
		maps["name"] = name
	}
	state := -1
	if arg := c.Query("state"); arg != "" {
		com.StrTo(arg).MustInt() //com包组合使用的
		maps["state"] = state
	}
	code := e.SUCCESS
	//gin的请求参数存在上下文中，可以在不同函数中分次处理
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func AddTag(c *gin.Context) {

}
func EditTag(c *gin.Context) {

}
func DeleteTag(c *gin.Context) {

}
