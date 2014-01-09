leopard
=======

A simple and fast mapping framework

How to use:

mapping.register("source name", datasource)

s := mapping.newSession("sourceName")

s.FindAll(&objSlice)

s.Find(id, &obj)

s.AsyncFind(id, &obj)

s.Query(sql).Get(&objSlice)

s.Query().Where(whereStr).OrderBy(orderStr).Limit(n).Join(&obj.r1....).JoinAll().Get(&objSlice, aSync)

s.Update(&o)

s.Update(&objSlice)

s.Insert(&o)

s.Insert(&objSlice)

s.Delete(&o)

s.Delete(&objSlice)

s.Delete(whereSql)

//所有CRUD操作支持异步，通过可选的aSync参数

s.BeginTrans()

s.Commit()

s.Rollback()

mapping.register(&schemaMapping)

mapping.register(&dbDriver)

mapping.register($gloalCacher)

mapping.register($logger)

//cache
//批量SQL设置
//执行速度必须快
//支持根据关联表字段进行查询




