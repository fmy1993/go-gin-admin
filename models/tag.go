package models

type Tag struct {
	Model
	//绑定json的语法,是由grom框架生效的
	Name       string `json:"name"` //tag的名字
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//直接把返回值作为参数声明在返回值域，可以在函数体直接使用
//map是gorm框架where的条件（一个sql的where）
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	//where后头是sql查询的条件，根据条件按分页查询
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags) //用指针避免大内存复制
	return                                                     //不用写 return tags了
}

// go 这种直接用返回值的语法应该不用指针的话可能也会返回一个副本
func GetTagTotal(maps interface{}) (count int) {
	//得到model对应的表，并做下一步操作
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
