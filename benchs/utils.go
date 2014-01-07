package benchs

import (
	"database/sql"
	"fmt"
	"os"
)

type Model struct {
	Id      int `qbs:"pk" sql:"pk"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

func NewModel() *Model {
	m := new(Model)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://beego.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

var (
	ORM_MULTI    int
	ORM_MAX_IDLE int
	ORM_MAX_CONN int
	ORM_SOURCE   string
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func wrapExecute(b *B, cbk func()) {
	b.StopTimer()
	defer b.StartTimer()
	cbk()
}

func initDB() {
	sqls := []string{
		"DROP TABLE IF EXISTS `model`",
		"CREATE TABLE `orm_bench`.`model` (" +
			"`id` int(11) NOT NULL AUTO_INCREMENT," +
			"`name` varchar(255) NOT NULL," +
			"`title` varchar(255) NOT NULL," +
			"`fax` varchar(255) NOT NULL," +
			"`web` varchar(255) NOT NULL," +
			"`age` int(11) NOT NULL," +
			"`right` tinyint(1) NOT NULL," +
			"`counter` bigint(20) NOT NULL," +
			"PRIMARY KEY (`id`)" +
			") ENGINE=`INNODB` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci",
	}

	DB, err := sql.Open("mysql", ORM_SOURCE)
	checkErr(err)
	defer DB.Close()

	err = DB.Ping()
	checkErr(err)

	for _, sql := range sqls {
		_, err = DB.Exec(sql)
		checkErr(err)
	}
}
