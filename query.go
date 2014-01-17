package orm

import {
	"reflect"
}

const (
	query = iota
	insert
	delete
	update
)

//表示一个数据库操作，可以是CRUD中任意类型
type dbOperation struct {
	optType int
	objType reflect.Type
	obj		interface{}    //如果是IDU操作，obj是要保持的对象
	rawSql  string		//原生SQL
	args    interface{}   //原生SQL或查询对象的参数
	where   string //Where(whereStr).OrderBy(orderStr).Limit(n).offset(n)
	orderby string
	limit   string
	offset  string
}
