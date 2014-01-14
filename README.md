leopard
=======

A simple and fast mapping framework

设计目标：效率高，接口简单，支持原生SQL，支持完善的ORM功能

How to use:

## 设置与注册数据库

### 设置

##### 开启和关闭Debug

    mapping.Debug(bool) //开启debug logs
    
#### 注册数据库

    mapping.RegisterDb(connectStr, dbType)
    
#### 注册数据库驱动

    mapping.RegisterDriver(driver, dbType)
    
#### 注册各种自定义模块：

##### Log处理实现

    mapping.RegisterModule(Logger) //注册自定义的Logger实现
    
##### 注册全局缓存实现

    mapping.RegisterModule(Cacher) //注册自定义的全局缓存实现
    
#### 注册自定义映射实现

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

### 复杂查询

一个查询函数处理查询列表或者查询单个对象，根据传入的对象引用来进行判断

s.Query().Raw(sql).Find(&objSlice) //原生sql

s.Query().Where(whereStr).OrderBy(orderStr).Limit(n).Find(&objSlice, aSync) //根据提供的Where，Orderby，Limit进行查询

### 对关联对象的查询

//默认不对关联对象进行Join查询

//如果需要对关联对象进行Join查询，使用Join(“refobj1"....)或者JoinAll()，JoinAll查询所有关联对象（只查询一层），程序将生成join sql来提高效率，Join的参数是关联对象的Field名称

s.Query().Where(whereStr).Join(“refobj1"....).JoinAll().Find(&objSlice, aSync) //join

//JoinAll()表示查询所有关联对象

s.FindRef(&obj, obj.refobj1, obj.refobj2...)  //重新获取查询对象的关联对象，只打算支持1层关联，如果关联对象已经存在，则跳过

s.FindAllRef(&obj)              //重新获取查询对象的关联对象，只支持1层关联，如果关联对象已经存在，则跳过

//查询条件中允许根据关联对象的字段值进行查询

s.Query().Where("Address.state='hunan'").Find(&user) //Address可以是Struct的关联字段名，也可以是关联的表名

s.Query().Where("HomeAddress.state='hunan' and CompanyAddresslstate='GuangDong'").Find(&user) //如果一个表有多个关联字段，则必须为Struct的关联字段名，因为存在多个相同的关联的表名


### 对异步查询的支持

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



### 事务支持

s.BeginTrans()

s.Commit()

s.Rollback()

### Cache

Cache支持一级和二级缓存，查询对象的顺序为一级缓存》二级缓存》数据库。一级缓存保存的是对象引用，所以任何对对象内容的修改都直接反映到一级缓存。而二级缓存要等对象进行CUD操作或提交事务时才更新

### 批量SQL支持

当ORM进行批量的SQL操作时，比如更新/删除/插入列表时，ORM将自动设置批量SQL执行

### 性能优化

#### 通过Join和JoinAll来支持Lazy Loading或者Eager Loading

- 如果需要在加载列表时同时加载关联的对象，可以使用Join或者JoinAll

- 如果一开始不需要加载关联对象，而是在需要时加载某个对象的关联对象，可以使用FindRef或者FinAlldRef

- 如果一开始不需要加载关联对象，但后面又需要加载整个列表的关联对象，也可以使用FindRef或者FinAlldRef，参数为列表对象（针对这个需求，ORM应该根据之前的SQL语句生成一条SQL,而不是N条以提高性能，具体实现需要关联列表与原来的查询语句，还需要判断列表是否发生变化，这个可以通过记录列表的ID list的hash值来进行判断）

- 如果一个对象存在多个one to Many关联，那同样不适合使用Eager Loading，可以先只Eager loading一个关联对象，其他的采用FindRef的方式来加载

#### 支持预编译

#### 支持一级缓存与二级缓存






