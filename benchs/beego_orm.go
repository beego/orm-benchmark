package benchs

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

var bo orm.Ormer

func init() {
	st := NewSuite("orm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, BeegoOrmInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, BeegoOrmInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, BeegoOrmUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, BeegoOrmRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, BeegoOrmReadSlice)

		orm.RegisterDataBase("default", "mysql", ORM_SOURCE, ORM_MAX_IDLE, ORM_MAX_CONN)
		orm.RegisterModel(new(Model))

		bo = orm.NewOrm()
	}
}

func BeegoOrmInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := bo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmInsertMulti(b *B) {
	var ms []*Model
	wrapExecute(b, func() {
		initDB()

		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := bo.InsertMulti(100, ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		bo.Insert(m)
	})

	for i := 0; i < b.N; i++ {
		if _, err := bo.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		bo.Insert(m)
	})

	for i := 0; i < b.N; i++ {
		if err := bo.Read(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BeegoOrmReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if _, err := bo.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if _, err := bo.QueryTable("model").Filter("id__gt", 0).Limit(100).All(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
