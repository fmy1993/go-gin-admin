/*
 * @Descripttion:
 * @version:
 * @Author: fmy1993
 * @Date: 2021-04-23 15:27:23
 * @LastEditors: fmy1993
 * @LastEditTime: 2021-04-25 16:23:07
 */
package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
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
/*
@ c *gin.Context will use gin to handle http parameter
*/

/**
 * @description:请求处理函数，调用model包里的DAO函数，逻辑在这里解耦
 * @test: http请求处理放在router包下，DAO函数放在models包
 * @param {*gin.Context} c
 * @return 请求处理函数不需要返回值，而是用http状态码反映
 * @author: fmy1993
 */
func GetTags(c *gin.Context) {
	// c.JSON(200, gin.H{"msg": "send by router"})
	name := c.Query("name")
	//go的map自带动态添加属性，而且直接赋值就行，就像无限长的hashtable
	//maps:sql的查询条件
	maps := make(map[string]interface{})
	//data:查询数据的集合
	data := make(map[string]interface{})
	// check http parameter is not null,
	//other check can be done in front end
	if name != "" { //把要查询的sql的条件存进maps，传给gorm，实现类似where name="fmy",并且可以有多个
		maps["name"] = name
	}
	state := -1
	if arg := c.Query("state"); arg != "" {
		com.StrTo(arg).MustInt() //com包组合使用的,用于格式转换
		maps["state"] = state
	}
	code := e.SUCCESS
	//gin的请求参数存在上下文中，可以在不同函数中分次处理
	// models是包名，包内任意函数都可以直接由包名取到(go没有类的概念了，包名当做类)
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func AddTag(c *gin.Context) {
	name := c.Query("name")
	//DefaultQuery就是如果url没有这个key,就默认用第二个参数的值,如果有就用请求参数的值
	// 这样即使后期借口有调整也不会出问题
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	valid := validation.Validation{}
	// 结果，状态都只是存在valid这个库里而已
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1") //这个函数本身是限制范围的，但是是int也就可以实现离散值
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

/**
 * @description:根据id更新name的值，同时更新modify的时间,同时将修改完的数据返回给前端
 * @test: 在router中用:设置url的参数，eg:/a/:b,name (gin上下文)c.Param("b")可以取到b的值
 * @param {*gin.Context} c
 * @return {*}
 * @author: fmy1993
 */
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")
	valid := validation.Validation{} // 直接得到一个默认所有属性都是空的初始值
	state := -1                      //因为go的特殊空值就是0，因此这里特殊的定义为-1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0和1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	//并没有对name做null检查，而是放在框架下
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTagById(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}

	}
	c.JSON(code, gin.H{
		"code": code,
		"msg:": e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

func DeleteTag(c *gin.Context) {
	//id := com.StrTo(c.Query("id")).MustInt()
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	//valid.Required(id, "id")
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		models.DeleteTagById(id)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
