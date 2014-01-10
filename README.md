leopard
=======

A simple and fast mapping framework

How to use:

## 设置与注册数据库

### 设置

#####Debug

    mapping.Debug(bool) //开启debug logs
    
####注册Log处理实现

    mapping.Loger(Logger) //注册自定义的Logger实现
    
####注册全局缓存实现

    mapping.GlobalCacher(Cacher) //注册自定义的全局缓存实现
    
####注册数据库

    mapping.RegisterDb(connectStr, dbType)
    
####注册数据库驱动

    mapping.RegisterDriver(driver, dbType)
    
####注册自定义映射实现

    mapping.RegitserMapping(Mapper)
    
### Session和对象CRUD操作


s := mapping.newSession("sourceName").using(dbName) //using可选，如果操作默认数据库无需using

s.Find(&objSlice)

s.Find(id, &obj)

s.Update(&o)

s.Update(&objSlice)

s.Insert(&o)

s.Insert(&objSlice)

s.Delete(&o)

s.Delete(&objSlice)

s.Delete(whereSql)

//所有CRUD操作支持异步，通过可选的aSync参数

###复杂查询

s.Query(sql).Get(&objSlice) //原生sql

s.Query().Where(whereStr).OrderBy(orderStr).Limit(n).Join(&obj.r1....).JoinAll().Get(&objSlice, aSync) //join

s.QueryRelater(&obj)  //查询对象的关联对象

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




