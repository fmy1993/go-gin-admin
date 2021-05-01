/*
 * @Descripttion: aricle相关CRUD DAO层,记得该层的函数会被controllor调用，从这点考虑返回值，取数据的要返回对应的数据
 * @version:
 * @Author: fmy1993
 * @Date: 2021-04-25 16:36:50
 * @LastEditors: fmy1993
 * @LastEditTime: 2021-04-27 11:33:43
 */
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model //这里model连名字都不需要，因为只是给gorm操作的
	// gorm会自动对类名+ID（必须都大写）的属性进行管理，将其拆分为 类名+ID,然后可以对这个类做关联
	TagID      int    `json:"tagid" gorm:"index"` //gorm框架声明索引
	Tag        Tag    `json:"tag"`                //Tag        int    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreateBy   string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// go中用函数名前的结构体指针来绑定函数与结构体的关系(作为没有类的弥补)
//
//

/**
 * @description:article类的钩子函数，新增数据之前调用grom实现的scope接口更新数据的字段
 * @test: 这里的字段一律都是实体类定义的字段名，并且只要在同一个包下就可以拿到
 * @param {*gorm.Scope} scope
 * @return {*}
 * @author: fmy1993
 */
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

/**
 * @description: 更新数据之前利用gorm实现的scope接口实现先更新这条数据的某个字段
 * @test: 使用函数名前指针制定对应的实体类 ， 在model里以 "gorm：primary_key:id"声明了公共字段的结构体，因此可以再scope里拿到
 * @param {*gorm.Scope} scope
 * @return {*}
 * @author: fmy1993
 */
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

/**
* @description:分页查询文章
* @test: gorm分页用法，先where,再页大小(limit)，再页数(offset)，最后find以参数形式将结果集传出给参数指针,
这里使用gorm 的preload关键字来避免对多条数据的分次多次关联对db的影响（思路是把id取出来，放在in做一次子查询）
article和标签的关系是一对多，对应会是有多条article数据存着不同的tag
* @param {int} pageSize
* @param {int} pageNum
* @param {interface{}} maps
* @return {*}
* @author: fmy1993
*/
func GetArticlePage(pageSize int, pageNum int, maps interface{}) (article []Article) {
	//db.Where(maps).Limit(pageSize).Offset(pageNum).Find(&article)
	db.Preload("Tag").Where(maps).Limit(pageSize).Offset(pageNum)
	return
}

/**
 * @description:得到总的记录数
 * @test: 得到总记录数： 先model选定数据类型，where确定查询条件，count统计总条数
 * @param {interface{}} maps
 * @return 总记录数 total
 * @author: fmy1993
 */
func GetArticleTotal(maps interface{}) (total int) {
	db.Model(&Article{}).Where(maps).Count(&total)
	return
}

/**
 * @description:根据id 得到文章
 * @test: gorm 关联的实现：1. StructID(ID大写) gorm会自动寻找表名为prifix_struct的类(数据库为小写) 2.使用Ralated关键字将对应的表数据存在对应的结构体属性中
 * @param {int} id
 * @return {*}
 * @author: fmy1993
 */
func GetArticleById(id int) (article Article) {

	//db.Where("id=?", id).Find(&article) //示例这里用的是first，能不能用find?不能，因为文章可能有多个属性
	db.Where("id=?", id).First(&article)
	db.Model(&article).Related(&article.Tag) //也就是把关联后的数据存入article的Tag属性对应的类中
	return
}
func ExistArticleById(id int) bool {
	var article Article
	//思路是取出数据传给结构体，用语言再处理
	db.Select("id").Where("id = ?", id).First(article)
	//id := db.Select("id").Where("id = ?", id).First(&Article{})
	//会把第一条数据赋值给Article{}，但是无法取到,因此只能用声明的形式，有一个显示的指针名
	return article.ID > 0
	// article.CreatedOn 定义在model类里的任意字段都是可以取到的

}

/**
 * @description:根据传入的map来更新数据
 * @test: 使用gorm的model关键字传入指针来确定对应的数据（表）,gorm的model其实就相当于得到一个表映射的类
 * @param {int} id 要修改数据的id
 * @param {interface{}} data 这里传入一个map ，就是可以表现 column=value的形式
 * @return {*}
 * @author: fmy1993
 */
func EditArticleById(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=?", id).Update(data) //delete有两个参数，而update只有一个.因此apdate需要model，而delete不需要
	return true
}

/**
 * @description: 根据id删除对应的数据
 * @test:
 * @param {int} id
 * @return {*}
 * @author: fmy1993
 */
func DeleteArticleById(id int) bool {
	db.Where("id=?", id).Delete(&Article{})
	return true
}

// func AddArticle(title string,content string,tagid int,) bool {

// }
/**
 * @description:新增一条数据
 * @test: interface==>别的类型 interface.(类型)
 * @param {map[string]interface{}} data 新增的字段值存储在maps里纯如
 * @return {*}
 * @author: fmy1993
 */
func AddArticle(data map[string]interface{}) bool {
	//结构体声明时不要加,号，但是初始化时要
	db.Create(&Article{ //新增的时候不用管model，但是对应的公共字段就需要利用gorm框架的钩子函数来赋值

		Title:    data["title"].(string), //interface{}的特殊的类型转换
		Content:  data["content"].(string),
		Desc:     data["desc"].(string),
		CreateBy: data["created_by"].(string),
		TagID:    data["tag_id"].(int),
		State:    data["state"].(int),
	})
	return true
}
