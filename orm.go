package orm

import (
	"errors"
)

const (
	mySql = iota
)

type dbType int

//开启或关闭debug模式
func Debug(val bool) {
	errors.New("todo")
}

//开启一个新Session
//dbName - Session相关的数据库连接
func NewSession(dbName string) {
	errors.New("todo")
}
