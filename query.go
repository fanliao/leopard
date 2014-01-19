package orm

import {
	"reflect"
    "unsafe"
    "errors"
}

const (
	query = iota
	insert
	delete
	update
)

type where struct{
    str string
    args interface{}
}

//表示一个数据库操作，可以是CRUD中任意类型
type dbOperation struct {
	optType int
	objType reflect.Type
	obj		interface{}    //如果是IDU操作，obj是要保存的对象
	rawSql  string		//原生SQL
	args    interface{}   //原生SQL或查询对象的参数
	where   //Where(whereStr).OrderBy(orderStr).Limit(n).offset(n)
	orderBy string
	limit   int
	offset  int
}


func (this dbOperation) getIterator() func()(unsafe.Pointer, ok, error){
    var i int = 0
    var count int
    oneObj bool
    
    if reflect.ValueOf(obj).Elem().Type().Kind() == reflect.Struct {
        oneObj = true;
    } else if reflect.ValueOf(obj).Elem().Type().Kind() == reflect.Slice{
        count = 1;
        oneObj = false;
        s := reflect.ValueOf(obj)
        count = s.Len()
    } else {
        return nil errors.New("operation object type is incorrect")
    }
    return func()( result unsafe.Pointer, ok bool, err error){
        if oneObj{
            if i == 0 {
                return InterfaceToPtr(obj), true, nil
            } else {
                return nil, false, nil
            }
        } else {
            if i < count {
                result = InterfaceToPtr(s.Index(i))
                err = nil
                ok = true
                i++
                return
            }
        }
        return 
    }, nil
}

//一个查询所有对象的操作，obj应该是一个对象的Slice
func NewQueryAll(obj interface{}) dbOperation{
    return nil
}

//一个空的查询操作
func NewQuery() dbOperation{
    return nil
}

//新增操作，可以是新增一个对象（obj是1个指针）或多个对象（obj是1个Slice）
func NewInsert(obj interface{})dbOperation{
    typ, err := GetStructType(obj)
    return dbOperation{insert, typ, obj}
}

//修改操作，可以是修改一个对象（obj是1个指针）或多个对象（obj是1个Slice）
func NewUpdate(obj interface{})dbOperation{
    typ, err := GetStructType(obj)
    return dbOperation{update, typ, obj}
}

//删除操作，可以是删除一个对象（obj是1个指针）或多个对象（obj是1个Slice）
func NewDelete(obj interface{})dbOperation{
    typ, err := GetStructType(obj)
    return dbOperation{delete, typ, obj}
}

//.Where(whereStr).OrderBy(orderStr).Limit(n).Find
func (this dbOperation) Where(whereStr string, args interface{}...) dbOperation{
    this.where = where{whereStr, args}
    return this
}

func (this dbOperation) OrderBy(orderBy string) dbOperation{
    this.orderBy = orderBy
    return this
}

func (this dbOperation) Limit(limit int) dbOperation{
    this.limit = limit
    return this
}

func (this dbOperation) Offset(offset int) dbOperation{
    this.offset = offset
    return this
}

func (this dbOperation) Raw(raw string) dbOperation{
    this.rawSql = raw
    return this
}

//如果允许级联操作，那么Expand函数将返回一个代表所有DB操作的slice
func (this dbOperation) Expand() []dbOperation{
    return nil
}
