leopard
=======

A simple and fast mapping framework

设计目标：效率高，接口简单，支持原生SQL，支持完善的ORM功能

How to use:

## 设置与注册数据库

### 设置

#####开启和关闭Debug

    mapping.Debug(bool) //开启debug logs
    
####注册数据库

    mapping.RegisterDb(connectStr, dbType)
    
####注册数据库驱动

    mapping.RegisterDriver(driver, dbType)
    
####注册各种自定义模块：

#####Log处理实现

    mapping.RegisterModule(Logger) //注册自定义的Logger实现
    
#####注册全局缓存实现

    mapping.RegisterModule(Cacher) //注册自定义的全局缓存实现
    
####注册自定义映射实现

    mapping.RegitserModule(Mapper) //注册自定义的结构名-表映射和结构字段名-数据库字段名映射处理模块
    
### Session和基本的对象CRUD操作

所有查询根据参数来判断要返回对象列表，单个对象还是Map列表

s := mapping.newSession("sourceName").using(dbName) //using可选，如果操作默认数据库无需using

s.Find(&objSlice)

s.Find(id, &obj)    //如果根据一个值Find一个对象，则这个值必须是主键值

s.Update(&o)

s.Update(&objSlice)

s.Insert(&o)

s.Insert(&objSlice)

s.Delete(&o)

s.Delete(&objSlice)

s.Delete(whereSql)

//所有CRUD操作支持异步，通过可选的aSync参数

###复杂查询

一个查询函数处理查询列表或者查询单个对象，根据传入的对象引用来进行判断

s.Query(sql).Find(&objSlice) //原生sql

s.Query().Where(whereStr).OrderBy(orderStr).Limit(n).Find(&objSlice, aSync) //根据提供的Where，Orderby，Limit进行查询

###对关联对象的查询

//Join查询指定的关联对象，JoinAll查询所有关联对象（只查询一层），程序将生成join sql来提高效率

s.Query().Where(whereStr).Join(“refobj1"....).JoinAll().Find(&objSlice, aSync) //join

//JoinAll()表示查询所有关联对象

s.FindRef(&obj)  //查询对象的关联对象，只打算支持1层关联

###对异步查询的支持

所有的Find与CRUD函数都支持async参数

//主键策略支持数据库自增或ORM库自动生成UID

//映射范例：

    type User struct{
	    Id int      `orm:auto_key;col(id);`
		Name string  `orm:col(NAME);size(30);type(varchar);notnull`
		Roles []*Role  `orm:m2m(table:User_Role,key:user_id,mKey:role_id)`
		Depart *Depart  `orm:one(depart_id);`
		CreateDate time.Time  `orm:auto_now`
		Typ  string   `orm:default(1)`
	}
	
	type Depart struct{
	    Id int            `orm:auto_uid;pk`
		Name string       
		Users []*User     `orm:many(depart_id)`
	}
	
	type Role struct{
	    Id int
		Name string
		Users []*User
		CreateDate time.Time
		Ptr id           `orm:no`
	}



//事务支持

s.BeginTrans()

s.Commit()

s.Rollback()

//注册组件

mapping.register(&schemaMapping)

mapping.register(&dbDriver)

mapping.register($gloalCacher)

mapping.register($logger)

//cache
//批量SQL设置
//执行速度必须快
//支持根据关联表字段进行查询




