/*
 * @Descripttion:DAO与后端交互的逻辑
 * @version:
 * @Author: fmy1993
 * @Date: 2021-04-23 17:02:57
 * @LastEditors: fmy1993
 * @LastEditTime: 2021-04-26 10:06:44
 */
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	//绑定json的语法,是由grom框架生效的
	Name       string `json:"name"` //tag的名字
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

/**
 * @description: 分页查询tags列表
 * @test: test content
 * @param {int} pageNum 分页的页数
 * @param {int} pageSize 分页的页面大小
 * @param {interface{}} maps 是 sql where的条件,以map形式传给gorm
 * @return tags gorm查询的结果集
 * @author: fmy1993
 */
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) { //直接把返回值作为参数声明在返回值域，可以在函数体直接使用
	//where后头是sql查询的条件，根据条件按分页查询 //map是gorm框架where的条件（一个sql的where）
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags) //用指针避免大内存复制
	// gorm一般都是把最后结果直接传给指针变量出去
	return //不用写 return tags了,但是实际上还是返回了tags
}

// go 这种直接用返回值的语法应该不用指针的话可能也会返回一个副本
/**
 * @description: 得到总的记录数，就是总的记录数，显示在前台页面
 * @test: test content
 * @param {interface{}} maps gorm sql where的条件
 * @return {*}
 * @author: fmy1993
 */
func GetTagTotal(maps interface{}) (count int) {
	//得到model对应的表，并做下一步操作
	// &Tag{} 只是一个空的结构体指针给框架
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

/**
 * @description: 根据标签名检查标签是否存在,在更新操作之前，用来先做检查数据是否存在使程序健壮
 * @test: 根据tag name查询数据库，根据id是否大于0来判断是否存在数据
 * @param {string} name tag的名字
 * @return bool
 * @author: fmy1993
 */
func ExistTagByName(name string) bool {
	var tag Tag
	// where 就是把第二个参数放到第一个参数?的位置的一个sql表达式
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID > 0 // go int 空值就是0
}

/**
 * @description:根据标签名id检查标签是否存在
 * @test: test content
 * @param {int} id
 * @return {*}
 * @author: fmy1993
 */
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?", id).First(&tag)
	return tag.ID > 0 //判断有记录的一个写法,虽然tag结构体没有id，但是在model声明的结构体的字段依然能取到
}
func DeleteTagById(id int) bool {
	db.Where("id=?", id).Delete(&Tag{})
	return true
}

/**
 * @description:根据id更新tag的任意字段
 * @test: 可以一次更新多个字段的值
 * @param {int} id
 * @param {interface{}} data 传入一个map，里头的key一定要和数据库字段名对上
 * @return {*}
 * @author: fmy1993
 */
func EditTagById(id int, data interface{}) bool {
	//数据更新用model
	db.Model(&Tag{}).Where("id=?", id).Update(data)
	return true
}

/**
 * @description:根据参数执行一条insert语句
 * @test: gorm框架新增数据直接传入带值的结构体即可
 * @param {string} name
 * @param {int} state
 * @param {string} createdBy
 * @return {*}
 * @author: fmy1993
 */
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

/**
 * @description:指定表新增数据之前设置记录时间的字段值，值为uinx开始时间(1970)到现在的秒数
 * @test: 函数名前的指针参数限制了函数的调用者，因此也就限制了表,函数会在结构体指针对应的表新增数据之前被调用
 * @param {*gorm.Scope} scope
 * @return {*}
 * @author: fmy1993
 */
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

/**
 * @description:
 * @test: test content
 * @param {*gorm.Scope} scope 是一个获取当前db操作的上下文
 * @return {*}
 * @author: fmy1993
 */
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
