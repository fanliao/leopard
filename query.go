package orm

const()
    query = iota
    insert
    delete
    update
)
//表示一个数据库操作，可以是CRUD中任意类型
type Query struct{
    
}