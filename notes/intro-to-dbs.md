# 数据库简介

我们的应用程序无法持久保存数据。不是很有趣。

从技术上讲，我们几乎可以使用任何东西来持久化数据。我们可以简单地在代码中打开一个文本文件并向其中写入数据。

从长远来看，文本文件不能很好地工作。

+ 规模不大
+ 条件是可能的
+ 再加上更多的问题

我们需要发明各种机制来使它们发挥作用，而我们基本上就是在重新发明数据库

我们将使用现有的选项，而不是创建自己的数据库。

有许多类型的数据库，都是为不同的情况设计的：

1. 关系数据库-Postgresql MYSQL
2. 文档存储-Mongodb
3. 图形数据库-Dgraph Neo4j
4. 密钥/值存储-Boltdb等

每个数据库都有优点和缺点。如果一个数据库在某件事上很出色，那么它就是在其他地方制造一个缺陷。因此，大多数大公司将各种数据库用于不同的任务

我们将使用Postgresql，它是



