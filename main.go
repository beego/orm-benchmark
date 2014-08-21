package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"github.com/nashtsai/orm-benchmark/benchs"

	_ "github.com/go-sql-driver/mysql"
)

type ListOpts []string

func (opts *ListOpts) String() string {
	return fmt.Sprint(*opts)
}

func (opts *ListOpts) Set(value string) error {
	if value == "all" || strings.Index(" "+strings.Join(benchs.BrandNames, " ")+" ", " "+value+" ") != -1 {
	} else {
		return fmt.Errorf("wrong run name %s", value)
	}
	*opts = append(*opts, value)
	return nil
}

func (opts ListOpts) Shuffle() {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(opts); i++ {
		a := rd.Intn(len(opts))
		b := rd.Intn(len(opts))
		opts[a], opts[b] = opts[b], opts[a]
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var orms ListOpts
	flag.IntVar(&benchs.ORM_MAX_IDLE, "max_idle", 200, "max idle conns")
	flag.IntVar(&benchs.ORM_MAX_CONN, "max_conn", 200, "max open conns")
	flag.StringVar(&benchs.ORM_SOURCE, "source", "root:@/orm_bench?charset=utf8", "mysql dsn source")
	flag.IntVar(&benchs.ORM_MULTI, "multi", 1, "base query nums x multi")
	flag.Var(&orms, "orm", "orm name: all, "+strings.Join(benchs.BrandNames, ", "))
	flag.Parse()

	var all bool

	if len(orms) == 0 {
		all = true
	} else {
		for _, n := range orms {
			if n == "all" {
				all = true
			}
		}
	}

	if all {
		orms = ListOpts(benchs.BrandNames)
	}

	orms.Shuffle()

	for _, n := range orms {
		fmt.Println(n)
		benchs.RunBenchmark(n)
	}

	fmt.Println("\nReports: \n")
	fmt.Print(benchs.MakeReport())

}
