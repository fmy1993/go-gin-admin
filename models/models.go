/*
 * @Descripttion: 连接数据库的逻辑，以及定义每张表通用的字段
 * @version:
 * @Author: fmy1993
 * @Date: 2021-04-07 16:04:11
 * @LastEditors: fmy1993
 * @LastEditTime: 2021-04-26 09:11:40
 */

package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/fmy1993/go-gin-admain/pkg/setting"
)

/*
既然是db orm框架，涉及到连接db的信息，那么就可以直接从写好的setting包里读取需要的配置信息
*/
// gorm包的一个上下文
var db *gorm.DB

//每个实体类都会用到的字段，作为引用放在其他结构体内部，抽象共同，用了特殊的语法声明了作为id是作为gorm框架的主键
//那么在 scope *gorm.Scope 里就可以拿到这个结构体里的其他字段(会默认是公共的字段)
type Model struct {
	//主键有特殊语法声明，并且使用gorm包
	ID         int `gorm:"primary_key" json:"id"` //注意公共字段的声明
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// 直接写在函数的构造函数中，没有写get函数,那么其他包直接取到db指针就行了
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	//拿到ini包的context
	// ini里的这个库就是输入key拿到value
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		//%v,只输出第三个参数结构体的value值（不输出属性值）,
		// fatal 表示严重错误，执行完直接退出，就直接在这里中断了
		log.Fatal(2, "Fail to get section 'database':%v", err)
	}
	// 从配置文件中获取基本DB的配置信息，一定要导入setting包
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	//根据配置信息开启数据库连接
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}
	// 拿到这张表的上下文，传入表名参数，返回带前缀的表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	//在Gorm中，表名是结构体名的复数形式，列名是字段名的蛇形小写。

	//即，如果有一个user表，那么如果你定义的结构体名为：User，gorm会默认表名为users而不是user。
	// 全局禁用表名复数
	//这样的话，表名默认即为结构体的首字母小写形式。
	//gorm是go的orm框架，定义了结构体后，默认表名是结构体小写复数，
	//映射的表的字段名是蛇形命名法，eg:CityName-->city_name
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	//defer 的用法：保证最后这个连接一定会被关闭
	//defer 关键字用于修饰一个函数或者方法，使得该函数或者方法在返回前才会执行，也就说被延迟，但又可以保证一定会执行。
	defer db.Close()
}
