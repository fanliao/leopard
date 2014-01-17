package orm

type QueryEngineer interface {
	Exec(dbOpt dbOperation) (result, error)
	//Sql(dbOpt dbOperation) (string, interface{}...)

}

type dbAdapter interface {
}

type queryer interface {
}

type Result interface {
}
