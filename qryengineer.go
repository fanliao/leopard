package orm

type qryEngineer struct {
}

type sql struct{
	sql string
	args []interface{}
}

//执行一个dbOperation对象并返回结果
func (this qryEngineer) Exec(dbOpt dbOperation) (result Result, err error) {
	sql, args, err := this.sql(dbOpt)
	result, err = this.execSql(sql, args...)
	return result, err
}

//生成dbOperation对应的SQL,1个dbOperation可以对应多条sql
//此处要考虑缓存
func (this qryEngineer) sql(dbOpt dbOperation) ([]sql, error) {

}

func (this qryEnginerr) execBatchSql(sqls []sql) （Result, error) {
}

//执行SQL
func (this qryEngineer) execSql(sql string, interface{}...) (Result, error){
	
}


