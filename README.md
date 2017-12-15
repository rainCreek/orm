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

```
C:\Users\80545>curl -d "username=ooo&departname=1" http://localhost:8080/service/userinfo
{
  "UID": 1,
  "UserName": "ooo",
  "DepartName": "1",
  "CreateAt": "2017-11-29T18:21:42.9758369+08:00"
}
C:\Users\80545>curl http://localhost:8080/service/userinfo?userid=1
{
  "UID": 1,
  "UserName": "ooo",
  "DepartName": "1",
  "CreateAt": "2017-11-29T18:21:42.9758369+08:00"
}
```


### ab测试


xorm：
```
D:\Apache\Apache24\bin>ab -n 1000 -c 100 http://localhost:8080/?userid=
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /?userid=
Document Length:        19 bytes

Concurrency Level:      100
Time taken for tests:   2.035 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      176000 bytes
HTML transferred:       19000 bytes
Requests per second:    435.76 [#/sec] (mean)
Time per request:       229.484 [ms] (mean)
Time per request:       2.295 [ms] (mean, across all concurrent requests)
Transfer rate:          74.90 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   10  29.0      0     162
Processing:     3  207  50.7    206     446
Waiting:        2  205  52.6    206     446
Total:          4  217  42.7    211     446

Percentage of the requests served within a certain time (ms)
  50%    211
  66%    221
  75%    237
  80%    261
  90%    278
  95%    301
  98%    312
  99%    317
 100%    446 (longest request)
```

database/sql

```
D:\Apache\Apache24\bin>ab -n 1000 -c 100 http://localhost:8080/?userid=
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /?userid=
Document Length:        19 bytes

Concurrency Level:      100
Time taken for tests:   1.564 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      176000 bytes
HTML transferred:       19000 bytes
Requests per second:    585.77 [#/sec] (mean)
Time per request:       169.374 [ms] (mean)
Time per request:       1.690 [ms] (mean, across all concurrent requests)
Transfer rate:          98.29 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   10  19.0      0     92
Processing:    13  114  50.7    106     446
Waiting:        5   51  23.5     46     157
Total:         24  157  42.7    161     336

Percentage of the requests served within a certain time (ms)
  50%    109
  66%    125
  75%    118
  80%    122
  90%    120
  95%    132
  98%    153
  99%    169
 100%    242 (longest request)
```


## 对比 database/sql 与 xorm 

由以上ab测试可以看出

| |database/sql|xorm
|--|:--:|:--:|
|编程效率|相比xorm编程效率较低 |封装了SQL语句，编程效率更高 |
|程序结构|提供数据库操作接口|形式上比较简洁 |
|服务性能|相比xorm花费的时间稍少 |对请求服务的时间会增加 |
