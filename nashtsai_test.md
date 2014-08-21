# ORM Benchmark

A benchmark to compare the performance of golang orm package.

## Results (2014-8-21)

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
         qbs:     6.59s       658525 ns/op    5200 B/op     85 allocs/op
        xorm:     6.85s       685079 ns/op    2356 B/op     63 allocs/op
         raw:     7.50s       750174 ns/op     330 B/op     10 allocs/op
         orm:     9.52s       951895 ns/op    1577 B/op     38 allocs/op
        hood:     9.83s       983200 ns/op   13457 B/op    193 allocs/op

    2500 times - MultiInsert 100 row
        xorm:    11.31s      4523096 ns/op  116315 B/op   1864 allocs/op
         raw:    12.10s      4841086 ns/op   50807 B/op    413 allocs/op
         orm:    12.96s      5184357 ns/op   88203 B/op   1245 allocs/op
         qbs:     Not support multi insert
        hood:     Not support multi insert

   10000 times - Update
         raw:     7.07s       707244 ns/op     344 B/op     10 allocs/op
        xorm:     7.97s       796539 ns/op    2760 B/op     83 allocs/op
         qbs:     8.60s       859930 ns/op    5195 B/op     85 allocs/op
         orm:     9.97s       996826 ns/op    1507 B/op     37 allocs/op
        hood:    10.75s      1075427 ns/op   13454 B/op    193 allocs/op

   20000 times - Read
         raw:     3.81s       190633 ns/op     924 B/op     23 allocs/op
         qbs:     4.34s       216836 ns/op    7591 B/op    136 allocs/op
        xorm:     5.89s       294564 ns/op    7214 B/op    177 allocs/op
        hood:     7.53s       376511 ns/op    4784 B/op     68 allocs/op
         orm:     7.65s       382452 ns/op    2801 B/op     75 allocs/op

   10000 times - MultiRead limit 100
         raw:     6.29s       629051 ns/op   32696 B/op    815 allocs/op
         orm:    17.05s      1704526 ns/op   97280 B/op   3286 allocs/op
         qbs:    19.65s      1965080 ns/op  222824 B/op   5217 allocs/op
        hood:    22.82s      2281711 ns/op  238446 B/op   5552 allocs/op
        xorm:    24.57s      2456610 ns/op  195990 B/op   6026 allocs/op
```

#### Sample 2
```
     10000 times - Insert
          xorm:     6.85s       685162 ns/op    2357 B/op     63 allocs/op
           qbs:     6.94s       693856 ns/op    5188 B/op     85 allocs/op
           raw:     7.20s       719666 ns/op     330 B/op     10 allocs/op
           orm:     8.40s       839692 ns/op    1585 B/op     38 allocs/op
          hood:     9.22s       921943 ns/op   13438 B/op    193 allocs/op

      2500 times - MultiInsert 100 row
          xorm:    10.36s      4143998 ns/op  116300 B/op   1864 allocs/op
           raw:    11.63s      4652087 ns/op   50798 B/op    413 allocs/op
           orm:    12.78s      5111329 ns/op   88295 B/op   1246 allocs/op
           qbs:     Not support multi insert
          hood:     Not support multi insert

     10000 times - Update
           raw:     7.61s       760759 ns/op     344 B/op     10 allocs/op
           qbs:     8.11s       810840 ns/op    5185 B/op     85 allocs/op
          hood:     9.32s       932424 ns/op   13440 B/op    193 allocs/op
          xorm:     9.52s       952424 ns/op    2761 B/op     83 allocs/op
           orm:    10.28s      1028152 ns/op    1512 B/op     37 allocs/op

     20000 times - Read
           raw:     3.74s       187029 ns/op     925 B/op     23 allocs/op
           qbs:     4.38s       219125 ns/op    7582 B/op    136 allocs/op
          xorm:     5.85s       292319 ns/op    7232 B/op    177 allocs/op
          hood:     7.48s       373878 ns/op    4795 B/op     68 allocs/op
           orm:     7.63s       381272 ns/op    2819 B/op     75 allocs/op

     10000 times - MultiRead limit 100
           raw:     6.50s       649588 ns/op   32715 B/op    815 allocs/op
           orm:    17.23s      1722633 ns/op   97728 B/op   3287 allocs/op
           qbs:    19.55s      1954671 ns/op  222608 B/op   5217 allocs/op
          hood:    22.76s      2275626 ns/op  239531 B/op   5551 allocs/op
          xorm:    24.71s      2471025 ns/op  195905 B/op   6026 allocs/op
```
