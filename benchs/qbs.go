package benchs

import (
	"fmt"

	"github.com/coocood/qbs"
)

var qo *qbs.Qbs

func init() {
	st := NewSuite("qbs")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, QbsInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, QbsInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, QbsUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, QbsRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, QbsReadSlice)

		qbs.Register("mysql", ORM_SOURCE, "model", qbs.NewMysql())
		qbs.ChangePoolSize(ORM_MAX_IDLE)
		qbs.SetConnectionLimit(ORM_MAX_CONN, true)

		qo, _ = qbs.GetQbs()
	}
}

func QbsInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := qo.Save(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func QbsInsertMulti(b *B) {
	panic(fmt.Errorf("Not support multi insert"))

	var ms []*Model
	wrapExecute(b, func() {
		initDB()

		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		if err := qo.BulkInsert(ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func QbsUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		qo.Save(m)
	})

	for i := 0; i < b.N; i++ {
		if _, err := qo.Save(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func QbsRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		qo.Save(m)
	})

	for i := 0; i < b.N; i++ {
		if err := qo.Find(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func QbsReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if _, err := qo.Save(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := qo.Where("id > ?", 0).Limit(100).FindAll(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
