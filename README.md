leopard
=======

A simple and fast mapping framework

How to use:
mapping.register("source name", datasource)

s := mapping.newSession("sourceName")

s.FindAll(&objSlice)

s.Find(id)

s.Query(sql)
