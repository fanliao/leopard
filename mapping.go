////映射范例：

//  type User struct{
//	    Id int      `orm:auto_key;col(id);`
//		Name string  `orm:col(NAME);size(30);type(varchar);notnull`
//		Roles []*Role  `orm:m2m(table:User_Role,key:user_id,mKey:role_id)`
//		Depart *Depart `orm:one(depart_id);`
//		CreateDate time.Time  `orm:auto_now`
//		Typ  string   `orm:default(1)`
//	}

//	type Depart struct{
//	    Id int            `orm:auto_uid;pk`
//		Name string
//		Users []*User     `orm:many(depart_id)`
//	}

//	type Role struct{
//	    Id int
//		Name string
//		Users []*User
//		CreateDate time.Time
//		Ptr id           `orm:no`
//	}

package orm

import (
	"reflect"
    "sync"
)

const (
	cascade_none = iota
	cascade_delete
	cascade_update
)

type tableMapping struct {
	typ         reflect.Type
	tableName   string
	pk          string
	propMapping map[string]interface{}
}

type propMapping struct {
	propName string
}

type colMapping struct {
	colName string
	size    int
	dbType  string
	null    bool
	propMapping
}

type oneMapping struct {
	refType    reflect.Type
	myColName  string
	refColName string
	cascade    int
	linkToMe   bool
	propMapping
}

type manyMapping struct {
	refType    reflect.Type
	myColName  string
	refColName string
	cascade    int
	linkToMe   bool
	propMapping
}

//`orm:m2m(table:User_Role,key:user_id,mKey:role_id)`
type m2mMapping struct {
	refType       reflect.Type //关联对象的类型
	myColName     string       //自身的数据库关联字段名
	refColName    string       //关联表的数据库关联字段名
	midTableName  string       //中间表的表名
	midMyColName  string       //与自身关联的中间表字段名
	midRefColName string       //与关联表关联的中间表字段名
	cascade       int          //级联设置
	linkToMe      bool         //关联对象是否也关联到自身
	propMapping                //映射的属性名
}

type mappings struct {
    ms map[reflect.Type] *tableMapping
    lock sync.RWMutex
}

func (this *mappings) get(typ reflect.Type) *tableMapping {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if val, ok := this.ms[typ]; ok {
		return val
	}
	return nil
}

func (this *mappings) set(typ reflect.Type, rw *tableMapping) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if val, ok := this.ms[typ]; !ok {
		this.ms[typ] = rw
	}
}


func init(){
    mappings = mappings{make(map[reflect.Type]*tableMapping), new(sync.RWMutex)}
}

func getMapping(typ reflect.Type) (*tableMapping, error) {
    mapping := mappings.get(type)
    if mapping != nil {
        return mapping
    } else {
        mapping = createMapping(typ)
        mappings.set(typ, mapping)
        return mapping
    }
    
}

func createMapping(typ reflect.Type) (*tableMapping, error) {
    return nil, nil
}