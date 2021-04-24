package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/fmy1993/go-gin-admain/pkg/setting"
)

// 默认[get] ?page=x
// 第一页返回0
func GetPage(c *gin.Context) int { //返回分页的int
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int() //这里似乎就是处理一下page的参数，似乎不一定非要用这个包
	if page > 0 {
		result = (page - 1) * setting.PageSize //就是返回分页在db表里的起始条数信息
	}
	return result
}
