package orm

//Session代表了一个数据库连接
type Session struct {
	dber dber //数据库适配器
	cacheMgr cacheMgrer //Session cache
}

//查找记录，obj必须为ptr，可以是单个对象的ptr或者一个slice的ptr。读取的数据填充到obj中
//如果是单个对象的ptr，则根据对象的主键Field的值读取记录，如果是slice的ptr，则读取所有的记录。如果单个对象又未设置主键Field的值，则返回error
func (this Session) Find(obj interface{}) error {
	//newQuery产生查询对象和查询参数对象，为了性能优化，查询对象和产生的SQL将进行缓存
	//查询对象和参数对象交由db接口执行，db接口将产生最终的SQL并得到结果集
	result := dber.exec(newQuery(...))  //queryer, args := newQuery(...)
	//mapping对象根据结果集生成返回的对象
	error := mapping.toStruct(result, obj)
	//处理缓存
	error := cacheMgr.cache(queryer, args, obj)
	return error
}

//新增记录，obj必须为ptr，可以是单个对象的ptr或者一个slice的ptr。
func (this Session) Insert(obj interface{}) error {
	//newInsert产生插入对象和插入参数对象，为了性能优化，插入对象和产生的SQL将进行缓存
	//插入对象和参数对象交由db接口执行，db接口将产生最终的SQL并得到结果集
	result := dber.exec(newInsert(...))  //queryer, args := newQuery(...)
	//如果有自增或者timestamp，mapping对象将使用返回值更新原对象
	error := mapping.toStruct(result, obj)
	//处理缓存
	error := cacheMgr.cache(obj)
	return error
}

//修改记录，obj必须为ptr，可以是单个对象的ptr或者一个slice的ptr。
func (this Session) Update(obj interface{}) error {
	//newUpdate产生更新对象和更新参数对象，为了性能优化，更新对象和产生的SQL将进行缓存
	//更新对象和参数对象交由db接口执行，db接口将产生最终的SQL并得到结果集
	result := dber.exec(newUpdate(...))  //queryer, args := newQuery(...)
	//如果有timestamp，mapping对象将使用返回值更新原对象
	error := mapping.toStruct(result, obj)
	//处理缓存
	error = cacheMgr.cache(obj)
	return error
}

//删除记录，obj必须为ptr，可以是单个对象的ptr或者一个slice的ptr。
func (this Session) Delete(obj interface{}) error {
	//newDelete产生删除对象和删除参数对象，为了性能优化，删除对象和产生的SQL将进行缓存
	//删除对象和参数对象交由db接口执行，db接口将产生最终的SQL并得到结果集
	result := dber.exec(newDelete(...))  //queryer, args := newQuery(...)
	//如果有timestamp，mapping对象将使用返回值更新原对象
	//error := mapping.toStruct(result, obj)
	//处理缓存
	error := cacheMgr.remove(obj)
	return nil
}

