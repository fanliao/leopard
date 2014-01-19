package orm

type qryEngineer struct {
}

type sql struct{
	sql string
	args []interface{}
    callback func()  //当insert或者update时可能需要返回自增字段或timestamp字段的值，此逻辑可以使用callback回调函数完成
}

//执行一个dbOperation对象并返回结果
//为了优化sql执行，可以采取下面的方法：
//1. 多个SQL使用多个异步执行
//2. 多个SQL可以合并到一个batch语句中执行
//3. 预编译
func (this qryEngineer) Exec(dbOpt dbOperation) (result Result, err error) {
	sqls, err := this.sql(dbOpt)
	result, err = this.execSql(sqls)
	return result, err
}

//生成dbOperation对应的SQL,1个dbOperation可以对应多条sql
//此处要考虑缓存
func (this qryEngineer) sql(dbOpt dbOperation) ([]sql, error) {
    switch dbOpt.optType{
        case query:
        case insert:
        case update:
        case delete:
            //create delete sql
            itr := dbOpt.getIterator()
            sql := 
            for next, ok, err := itr();ok {
                whereStr := 
            }
            
    }

}

func (this qryEnginerr) execBatchSql(sqls []sql) （Result, error) {
}

//执行SQL
func (this qryEngineer) execSql(sql string, interface{}...) (Result, error){
	
}

func (this qryEngineer) getDeleteSql() {
    
}

func 


