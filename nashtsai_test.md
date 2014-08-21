# ORM Benchmark

A benchmark to compare the performance of golang orm package.

## Results (2014-8-22)

### Environment

* MBP Intel Core i5 2.40GHz (2 cores)
* 8G RAM
* Mac OS X 10.9.4
* go version go1.3.1 darwin/amd64
* [Go-MySQL-Driver Latest](https://github.com/go-sql-driver/mysql)

### MySQL

* MySQL 5.6.13 Server for Mac OS X on x86_64

### ORMs

All package run in no-cache mode.

* [Beego ORM](http://beego.me/docs/mvc/model/overview.md) latest
* [xorm](https://github.com/go-xorm/xorm) latest
* [Hood](https://github.com/eaigner/hood) latest
* [Qbs](https://github.com/coocood/qbs) latest (Disabled stmt cache / [patch](https://gist.github.com/slene/8297019) / [full](https://gist.github.com/slene/8297565))

### Run

```go
go get github.com/nashtsai/orm-benchmark
cd $GOPATH/src/github.com/nashtsai/orm-benchmark
go build
./orm-benchmark -multi=5 -orm=all
```

### Reports

#### Sample 1
```
 10000 times - Insert
       raw:     6.20s       619586 ns/op     329 B/op     10 allocs/op
       qbs:     7.14s       714323 ns/op    5181 B/op     85 allocs/op
      xorm:     7.78s       778007 ns/op    2356 B/op     63 allocs/op
       orm:     8.69s       869307 ns/op    1573 B/op     38 allocs/op
      hood:     9.65s       964711 ns/op   13497 B/op    193 allocs/op

  2500 times - MultiInsert 100 row
       raw:     9.80s      3920504 ns/op   50725 B/op    413 allocs/op
      xorm:    11.28s      4513046 ns/op  116264 B/op   1864 allocs/op
       orm:    13.36s      5343020 ns/op   88042 B/op   1245 allocs/op
      hood:     Not support multi insert
       qbs:     Not support multi insert

 10000 times - Update
       qbs:     6.59s       659403 ns/op    5181 B/op     85 allocs/op
       raw:     7.14s       714484 ns/op     343 B/op     10 allocs/op
      xorm:     8.29s       829202 ns/op    2759 B/op     83 allocs/op
      hood:     9.29s       928678 ns/op   13477 B/op    193 allocs/op
       orm:     9.71s       970709 ns/op    1504 B/op     37 allocs/op

 20000 times - Read
       raw:     3.69s       184264 ns/op     921 B/op     23 allocs/op
       qbs:     4.35s       217527 ns/op    7568 B/op    136 allocs/op
      xorm:     5.82s       290913 ns/op    7186 B/op    177 allocs/op
      hood:     7.31s       365405 ns/op    4831 B/op     68 allocs/op
       orm:     7.56s       377792 ns/op    2772 B/op     75 allocs/op

 10000 times - MultiRead limit 100
      xorm:     5.29s       528933 ns/op   35646 B/op    765 allocs/op
       raw:     6.87s       687438 ns/op   32618 B/op    815 allocs/op
       orm:    17.07s      1706562 ns/op   96891 B/op   3286 allocs/op
       qbs:    19.64s      1963853 ns/op  221316 B/op   5216 allocs/op
      hood:    22.66s      2265802 ns/op  239176 B/op   5552 allocs/op
```

#### Sample 2
```
 10000 times - Insert
       raw:     6.39s       638549 ns/op     330 B/op     10 allocs/op
      xorm:     6.86s       686125 ns/op    2364 B/op     63 allocs/op
       qbs:     6.90s       689705 ns/op    5203 B/op     85 allocs/op
       orm:     8.79s       879231 ns/op    1574 B/op     38 allocs/op
      hood:     9.31s       930851 ns/op   13426 B/op    193 allocs/op

  2500 times - MultiInsert 100 row
      xorm:     7.50s      2999451 ns/op  116465 B/op   1864 allocs/op
       raw:     8.59s      3435754 ns/op   50833 B/op    413 allocs/op
       orm:    10.82s      4326730 ns/op   88097 B/op   1245 allocs/op
       qbs:     Not support multi insert
      hood:     Not support multi insert

 10000 times - Update
       qbs:     6.98s       697781 ns/op    5199 B/op     85 allocs/op
      xorm:     7.28s       727931 ns/op    2763 B/op     83 allocs/op
       raw:     8.11s       811301 ns/op     344 B/op     10 allocs/op
       orm:     9.45s       945375 ns/op    1505 B/op     37 allocs/op
      hood:     9.81s       981114 ns/op   13425 B/op    193 allocs/op

 20000 times - Read
       raw:     3.65s       182250 ns/op     923 B/op     23 allocs/op
       qbs:     4.61s       230311 ns/op    7600 B/op    136 allocs/op
      xorm:     5.61s       280429 ns/op    7262 B/op    177 allocs/op
      hood:     7.47s       373731 ns/op    4766 B/op     68 allocs/op
       orm:     8.10s       404900 ns/op    2784 B/op     75 allocs/op

 10000 times - MultiRead limit 100
      xorm:     5.74s       574265 ns/op   35584 B/op    765 allocs/op
       raw:     6.77s       677382 ns/op   32672 B/op    815 allocs/op
       orm:    17.02s      1701782 ns/op   97191 B/op   3286 allocs/op
       qbs:    19.39s      1938879 ns/op  222771 B/op   5218 allocs/op
      hood:    23.14s      2314262 ns/op  238526 B/op   5550 allocs/op

```
