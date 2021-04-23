package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

// 得到page具体的页数,默认都是发送页数的get请求
// 使用com包,只需要传入第几页就行，默认get请求
func GetPage(c *gin.Context) int { //返回分页的int
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int() //这里似乎就是处理一下page的参数，似乎不一定非要用这个包
	if page > 0 {
		result = (page - 1) * setting.PageSize //就是返回分页在db表里的起始条数信息
	}
	return result
}
