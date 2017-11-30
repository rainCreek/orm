# orm
xorm，一个简单而强大的 Go 语言 ORM 框架

xorm优势：

* **少依赖；**
xorm除了依赖github.com/go-xorm/core之外不依赖其它第三方库。
* **易使用；** 
通过连写操作，可以通过很少的语句完成数据库操作。 
* **功能全；**
支持缓存，事务，乐观锁，多种数据库支持，反转等等特性。 
* **开源化**


## 主要内容

主要是改写`userinfo-service`中的 `Save(u *UserInfo)`、`FindAll()`、`FindByID(id int)`函数。

### save
需要使用xorm进行**事务处理**。

> 当使用事务处理时，需要创建**Session对象**。在进行事物处理时，可以混用ORM方法和RAW方法

这里主要是进行插入操作。

```
session := engine.NewSession()
...
_, err = session.Insert(&u)
...
```

### FindAll
1. 可通过调用**engine.DBMetas()** 可以获取到数据库中所有的表，字段，索引的信息
2. 查询多条数据使用**Find方法**

这里选用了比较易懂的Find方法

> Find方法的第一个参数为slice的指针或Map指针

```
everyone := make([]UserInfo, 0)
err := engine.Find(&everyone)
...
```

### FindByID
传入一个主键字段的值，作为查询条件：
1. 查询单条数据使用**Get方法**，根据Id来获得单条数据
2. 直接执行一个SQL查询

这里选用了比较直观的Get方法

```
...
engine.Id(id).Get(&user)
...
```

相当于 `SELECT * FROM user Where id = id`

## 测试
### 运行mysql并建立好数据库


### 插入数据


### 查询数据


### ab测试



## 对比 database/sql 与 xorm 

由以上ab测试可以看出

| |database/sql|xorm
|--|:--:|:--:|
|编程效率|相比xorm编程效率较低 |封装了SQL语句，编程效率更高 |
|程序结构|提供数据库操作接口|形式上比较简洁 |
|服务性能|相比xorm花费的时间稍少 |对请求服务的时间会增加 |
