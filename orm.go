package orm

type dbType int

const(
	mySql = iota
)

import(
	"errors"
)

//开启或关闭debug模式
func Debug(val bool){
	errors.New("todo")
}

//开启一个新Session
//dbName - Session相关的数据库连接
func NewSession(var dbName){
	errors.New("todo")
}



