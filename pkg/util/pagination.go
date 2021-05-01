/*
 * @Descripttion:
 * @version:
 * @Author: fmy1993
 * @Date: 2021-04-07 15:44:46
 * @LastEditors: fmy1993
 * @LastEditTime: 2021-04-27 10:18:42
 */
package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/fmy1993/go-gin-admain/pkg/setting"
)

/**
 * @description:根据get请求的的 ？page=x参数得到对应页数的第一条数据的id
 * @test: 在内部处理page这个请求参数,PageSize写在配置文件中，再用ini库来实现IoC，只要修改配置文件的值就能改变程序中所有用到这个值的地方
 * @param {*gin.Context} c,注意get请求的参数一定要是 query
 * @return 返回分页的第一条数据id int
 * @author: fmy1993
 */
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int() //这里似乎就是处理一下page的参数，似乎不一定非要用这个包
	if page > 0 {
		result = (page - 1) * setting.PageSize //就是返回分页在db表里的起始条数信息
	}
	return result
}
