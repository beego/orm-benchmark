# ORM Benchmark

A benchmark to compare the performance of golang orm package.

## Results (2014-1-14)

### Environment

* MBP Intel Core i5 2.40GHz (2 cores)
* 8G RAM
* Mac OS X 10.9.1
* go version go1.3.1 darwin/amd64
* [Go-MySQL-Driver Latest](https://github.com/go-sql-driver/mysql)

### MySQL

* MySQL 5.6.13 Server for Mac OS X on x86_64

### ORMs

All package run in no-cache mode.

* [Beego ORM](http://beego.me/docs/mvc/model/overview.md) latest in branch [develop](https://github.com/astaxie/beego/tree/develop)
* [xorm](https://github.com/go-xorm/xorm) latest
* [Hood](https://github.com/eaigner/hood) latest
* [Qbs](https://github.com/coocood/qbs) latest (Disabled stmt cache / [patch](https://gist.github.com/slene/8297019) / [full](https://gist.github.com/slene/8297565))

### Run

```go
go get github.com/nashtsai/orm-benchmark
cd $GOPATH/src/github.com/nashtsai/orm-benchmark
go build
./orm-benchmark
```

### Reports

#### Sample 1
```
  2000 times - Insert
       qbs:     1.26s       627834 ns/op    5205 B/op    102 allocs/op
       orm:     1.53s       763963 ns/op    1513 B/op     35 allocs/op
      xorm:     1.56s       780115 ns/op    2341 B/op     63 allocs/op
      hood:     1.95s       975093 ns/op   14617 B/op    293 allocs/op

   500 times - MultiInsert 100 row
      xorm:     1.76s      3527580 ns/op  115263 B/op   1871 allocs/op
       orm:     2.03s      4066070 ns/op   85006 B/op    853 allocs/op
      hood:     Not support multi insert
       qbs:     Not support multi insert

  2000 times - Update
       qbs:     1.33s       664496 ns/op    5198 B/op    102 allocs/op
       orm:     1.67s       836731 ns/op    1463 B/op     34 allocs/op
      xorm:     2.19s      1095675 ns/op    2612 B/op     83 allocs/op
      hood:     2.33s      1166249 ns/op   14571 B/op    293 allocs/op

  4000 times - Read
       qbs:     0.85s       212266 ns/op    7573 B/op    162 allocs/op
      xorm:     1.07s       268716 ns/op    6435 B/op    185 allocs/op
       orm:     1.48s       371019 ns/op    2748 B/op     81 allocs/op
      hood:     2.07s       517772 ns/op    4976 B/op     96 allocs/op

  2000 times - MultiRead limit 100
       orm:     2.69s      1347041 ns/op   90253 B/op   3201 allocs/op
       qbs:     3.51s      1752612 ns/op  220549 B/op   6132 allocs/op
      xorm:     3.81s      1905392 ns/op  178893 B/op   6099 allocs/op
      hood:     7.22s      3609463 ns/op  253911 B/op   9226 allocs/op
```

#### Sample 2
```
  2000 times - Insert
       qbs:     1.27s       636110 ns/op    5216 B/op    102 allocs/op
      xorm:     1.28s       642299 ns/op    2337 B/op     63 allocs/op
      hood:     1.75s       874980 ns/op   14556 B/op    293 allocs/op
       orm:     1.85s       925578 ns/op    1550 B/op     35 allocs/op

   500 times - MultiInsert 100 row
      xorm:     2.16s      4325997 ns/op  115131 B/op   1870 allocs/op
       orm:     2.40s      4795201 ns/op   85220 B/op    854 allocs/op
       qbs:     Not support multi insert
      hood:     Not support multi insert

  2000 times - Update
       qbs:     1.26s       629532 ns/op    5214 B/op    102 allocs/op
      xorm:     1.45s       725277 ns/op    2609 B/op     83 allocs/op
       orm:     1.70s       849041 ns/op    1465 B/op     34 allocs/op
      hood:     1.77s       883198 ns/op   14544 B/op    293 allocs/op

  4000 times - Read
       qbs:     0.88s       219476 ns/op    7567 B/op    162 allocs/op
      xorm:     1.11s       276379 ns/op    6437 B/op    185 allocs/op
      hood:     1.48s       369093 ns/op    4968 B/op     96 allocs/op
       orm:     1.53s       382862 ns/op    2772 B/op     81 allocs/op

  2000 times - MultiRead limit 100
       orm:     2.71s      1353003 ns/op   90508 B/op   3201 allocs/op
       qbs:     3.80s      1897643 ns/op  220994 B/op   6133 allocs/op
      xorm:     3.88s      1941303 ns/op  178964 B/op   6100 allocs/op
      hood:     4.92s      2460267 ns/op  253970 B/op   9225 allocs/op
```
