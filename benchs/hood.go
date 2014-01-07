package benchs

import (
	"database/sql"
	"fmt"

	"github.com/eaigner/hood"
)

var ho *hood.Hood

func init() {
	st := NewSuite("hood")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, HoodInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, HoodInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, HoodUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, HoodRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, HoodReadSlice)

		db, _ := sql.Open("mysql", ORM_SOURCE)
		ho = hood.New(db, hood.NewMysql())
	}
}

func HoodInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := ho.Save(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func HoodInsertMulti(b *B) {
	panic(fmt.Errorf("Not support multi insert"))
}

func HoodUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		ho.Save(m)
	})

	for i := 0; i < b.N; i++ {
		if _, err := ho.Save(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func HoodRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		ho.Save(m)
	})

	for i := 0; i < b.N; i++ {
		var mds []Model
		if err := ho.Where("id", "=", m.Id).Find(&mds); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func HoodReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if _, err := ho.Save(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []Model
		if err := ho.Where("id", ">", 0).Limit(100).Find(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
