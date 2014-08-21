# ORM Benchmark

A benchmark to compare the performance of golang orm package.

## Results (2014-1-7)

### Environment

* Aliyun Intel Xeon E5-2630 2.30GHz (4 core)
* 8G RAM
* CentOS 6.5
* go version go1.2 linux/amd64
* [Go-MySQL-Driver Latest](https://github.com/go-sql-driver/mysql)

### MySQL

* MySQL 5.5.34 for Linux on x86_64
* MySQL-server-5.5.34-1.rhel5.x86_64.rpm
* Config in my.cnf

### ORMs

All package run in no-cache mode.

* [Beego ORM](http://beego.me/docs/mvc/model/overview.md) latest in branch [develop](https://github.com/astaxie/beego/tree/develop)
* [xorm](https://github.com/go-xorm/xorm) latest
* [Hood](https://github.com/eaigner/hood) latest
* [Qbs](https://github.com/coocood/qbs) latest (Disabled stmt cache / [patch](https://gist.github.com/slene/8297019) / [full](https://gist.github.com/slene/8297565))

### Run

```go
go get github.com/beego/orm-benchmark
orm-benchmark -multi=20 -orm=all
```

### Reports

#### Sample 1

```
40000 times - Insert
       orm:    11.63s       290706 ns/op    1515 B/op     35 allocs/op
       qbs:    12.92s       322982 ns/op    5846 B/op    111 allocs/op
      xorm:    13.21s       330361 ns/op    2176 B/op     53 allocs/op
      hood:    14.26s       356591 ns/op   14592 B/op    293 allocs/op

 10000 times - MultiInsert 100 row
       orm:    22.17s      2217284 ns/op   85061 B/op    854 allocs/op
      xorm:    30.17s      3016717 ns/op  112572 B/op   1877 allocs/op
      hood:     Not support multi insert
       qbs:     Not support multi insert

 40000 times - Update
       orm:    12.77s       319335 ns/op    1463 B/op     34 allocs/op
       qbs:    12.91s       322643 ns/op    5843 B/op    111 allocs/op
      xorm:    12.93s       323136 ns/op    2500 B/op     73 allocs/op
      hood:    14.53s       363331 ns/op   14578 B/op    293 allocs/op

 80000 times - Read
       orm:    19.57s       244639 ns/op    2762 B/op     81 allocs/op
      hood:    19.76s       246961 ns/op    4981 B/op     96 allocs/op
       qbs:    20.44s       255498 ns/op    8609 B/op    180 allocs/op
      xorm:    27.35s       341934 ns/op    6807 B/op    193 allocs/op

 40000 times - MultiRead limit 100
       orm:    39.32s       982898 ns/op   90386 B/op   3201 allocs/op
       qbs:    62.58s      1564475 ns/op  221787 B/op   6151 allocs/op
      xorm:    83.19s      2079663 ns/op  189770 B/op   6816 allocs/op
      hood:    87.64s      2191093 ns/op  254457 B/op   9228 allocs/op
```

#### Sample 2

```
 40000 times - Insert
       orm:    11.43s       285684 ns/op    1511 B/op     35 allocs/op
      xorm:    12.32s       308044 ns/op    2182 B/op     53 allocs/op
       qbs:    12.94s       323575 ns/op    5849 B/op    111 allocs/op
      hood:    14.45s       361222 ns/op   14563 B/op    293 allocs/op

 10000 times - MultiInsert 100 row
       orm:    22.37s      2236512 ns/op   84921 B/op    853 allocs/op
      xorm:    25.49s      2549191 ns/op  112714 B/op   1878 allocs/op
      hood:     Not support multi insert
       qbs:     Not support multi insert

 40000 times - Update
       orm:    12.95s       323640 ns/op    1460 B/op     34 allocs/op
       qbs:    13.02s       325452 ns/op    5850 B/op    111 allocs/op
      xorm:    13.13s       328245 ns/op    2503 B/op     73 allocs/op
      hood:    14.70s       367445 ns/op   14559 B/op    293 allocs/op

 80000 times - Read
       orm:    19.36s       241961 ns/op    2744 B/op     81 allocs/op
      hood:    19.76s       246971 ns/op    4958 B/op     96 allocs/op
       qbs:    20.47s       255905 ns/op    8628 B/op    180 allocs/op
      xorm:    26.38s       329728 ns/op    6820 B/op    193 allocs/op

 40000 times - MultiRead limit 100
       orm:    39.48s       987091 ns/op   90160 B/op   3200 allocs/op
       qbs:    63.17s      1579254 ns/op  221992 B/op   6152 allocs/op
      xorm:    83.58s      2089464 ns/op  189706 B/op   6815 allocs/op
      hood:    91.41s      2285256 ns/op  254087 B/op   9226 allocs/op
```

#### Sample 3
```
 40000 times - Insert
       orm:    11.95s       298654 ns/op    1512 B/op     35 allocs/op
      xorm:    12.47s       311666 ns/op    2179 B/op     53 allocs/op
       qbs:    13.49s       337224 ns/op    5849 B/op    111 allocs/op
      hood:    15.00s       374967 ns/op   14573 B/op    293 allocs/op

 10000 times - MultiInsert 100 row
       orm:    22.95s      2294502 ns/op   84920 B/op    853 allocs/op
      xorm:    26.01s      2600964 ns/op  112579 B/op   1877 allocs/op
       qbs:     Not support multi insert
      hood:     Not support multi insert

 40000 times - Update
       orm:    13.17s       329226 ns/op    1461 B/op     34 allocs/op
       qbs:    13.27s       331845 ns/op    5847 B/op    111 allocs/op
      xorm:    13.64s       340992 ns/op    2502 B/op     73 allocs/op
      hood:    14.79s       369635 ns/op   14576 B/op    293 allocs/op

 80000 times - Read
       orm:    20.10s       251271 ns/op    2738 B/op     81 allocs/op
      hood:    20.45s       255610 ns/op    4960 B/op     96 allocs/op
       qbs:    21.13s       264091 ns/op    8651 B/op    180 allocs/op
      xorm:    27.60s       345010 ns/op    6801 B/op    193 allocs/op

 40000 times - MultiRead limit 100
       orm:    40.36s      1008896 ns/op   90137 B/op   3200 allocs/op
       qbs:    64.53s      1613287 ns/op  222005 B/op   6152 allocs/op
      xorm:    81.46s      2036502 ns/op  189573 B/op   6815 allocs/op
      hood:    90.95s      2273773 ns/op  253770 B/op   9225 allocs/op
```


### Contact

Maintain by [slene](https://github.com/slene)