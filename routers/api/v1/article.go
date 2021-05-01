/*
 * @Descripttion:
 * @version:
 * @Author: fmy1993
 * @Date: 2021-04-30 17:10:36
 * @LastEditors: fmy1993
 * @LastEditTime: 2021-05-01 08:46:03
 */

package v1

import (
	"log"

	"github.com/astaxie/beego/validation"
	"github.com/fmy1993/go-gin-admain/models"
	"github.com/fmy1993/go-gin-admain/pkg/error"
	"github.com/fmy1993/go-gin-admain/pkg/setting"
	"github.com/fmy1993/go-gin-admain/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetArticleList(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	valid := validation.Validation{}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		//data["state"] = state
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	tagId := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tagId"] = tagId

		valid.Min(tagId, 1, "tagId").Message("标签ID必须大于0")
	}
	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		code = error.SUCCESS
		data["articles"] = models.GetArticlePage(setting.PageSize, util.GetPage(c), maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		// valid这个框架实现把错误存储起来，实际上也是一个k,v的集合
		for _, err := range valid.Errors {
			//log.Println()
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})

}

//GetArticleById
func GetArticleById(c *gin.Context) {

	data := make(map[string]interface{})
	//maps := make(map[string]interface{})

	valid := validation.Validation{}

	code := error.INVALID_PARAMS

	if arg := c.Query("id"); arg != "" {
		code = error.SUCCESS
		id := com.StrTo(arg).MustInt()
		valid.Min(id, 1, "id").Message("id必须>0")
		data["article"] = models.GetArticleById(id)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})

}
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}
	code := error.INVALID_PARAMS
	id := com.StrTo(c.Query("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	content := c.Query("content")
	desc := c.Query("desc")
	modifiedBy := c.Query("modified_by")

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	//code := error.INVALID_PARAMS
	data := make(map[string]interface{})
	if !valid.HasErrors() {
		code = error.SUCCESS
		if models.ExistArticleById(id) {//检查文章是否存在
			if models.ExistTagByID(tagid) {//检查tag是否存在
				data["tag_id"]=tagId
				if title!="" {
					data["title"]=title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				data["modified_by"] = modifiedBy
				models.EditArticleById(id,data)
			}else{
				code = error.ERROR_NOT_EXIST_TAG
			}
		}else{
			code = error.ERROR_NOT_EXIIST_ARTICLE
		}
		c.JSON(code,gin.H{
			"code":code,
			"msg":error.GetMsg(code),
			"data":make(map[string]interface{}),
		})

}

func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	data := make(map[string]interface{})
	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		data["tag_id"] = tagId
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["created_by"] = createdBy
		data["state"] = state
		code = error.SUCCESS
		models.AddArticle(data)

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

/*
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}
	code := error.INVALID_PARAMS
	id := com.StrTo(c.Query("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	content := c.Query("content")
	desc := c.Query("desc")
	modifiedBy := c.Query("modified_by")

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	//code := error.INVALID_PARAMS
	data := make(map[string]interface{})
	if !valid.HasErrors() {
		code = error.SUCCESS
		if models.ExistArticleById(id) {//检查文章是否存在
			if models.ExistTagByID(tagid) {//检查tag是否存在
				data["tag_id"]=tagId
				if title!="" {
					data["title"]=title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				data["modified_by"] = modifiedBy
				models.EditArticleById(id,data)
			}else{
				code = error.ERROR_NOT_EXIST_TAG
			}
		}else{
			code = error.ERROR_NOT_EXIIST_ARTICLE
		}
		c.JSON(code,gin.H{
			"code":code,
			"msg":error.GetMsg(code),
			"data":make(map[string]interface{}),
		})

}
*/

/*
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	code = error.INVALID_PARAMS
	valid.Min(id,1,"id").Message("id必须大于0")
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			models.DeleteArticleById(id)
			code=error.SUCCESS
		}else{
			code=error.ERROR_NOT_EXIIST_ARTICLE
		}
	}else{
		for _, err := range valid.Errors {
            log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
        }
	}
	c.JSON(code,gin.H{
		"code":code,
		"msg":error.GetMsg(code),
		"data":make(map[string]interface{}),

	})
}
*/
//