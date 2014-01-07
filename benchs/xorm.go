package benchs

import (
	"fmt"

	"github.com/lunny/xorm"
)

var xo *xorm.Session

func init() {
	st := NewSuite("xorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, XormInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, XormInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, XormUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, XormRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, XormReadSlice)

		engine, _ := xorm.NewEngine("mysql", ORM_SOURCE)

		engine.SetMaxIdleConns(ORM_MAX_IDLE)
		engine.SetMaxConns(ORM_MAX_CONN)

		xo = engine.NewSession()
	}
}

func XormInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := xo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormInsertMulti(b *B) {
	var ms []*Model
	wrapExecute(b, func() {
		initDB()

		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := xo.InsertMulti(&ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		xo.Insert(m)
	})

	for i := 0; i < b.N; i++ {
		if _, err := xo.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		xo.Insert(m)
	})

	for i := 0; i < b.N; i++ {
		if _, err := xo.Get(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if _, err := xo.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := xo.Where("id > ?", 0).Limit(100).Find(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
