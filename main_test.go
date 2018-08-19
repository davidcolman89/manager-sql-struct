package sqlstruct

import (
	"database/sql"
	"testing"
	"time"
	"github.com/go-sql-driver/mysql"
)

type Sql struct {
	Name     sql.NullString
	Age      sql.NullInt64
	Balance  sql.NullFloat64
	Valid    sql.NullBool
	Name2    string
	Balance2 float64
	Balance3 float32
	Age2     int
	Age3     int8
	Age4     int16
	Age5     int32
	Age6     int32
	Age7     int64
	Age8     uint
	Age9     uint8
	Age10    uint16
	Age11    uint32
	Age12    uint64
	Date     mysql.NullTime
}

type Model struct {
	Name     string
	Age      int64
	Balance  float64
	Valid    bool
	Name2    string
	Balance2 float64
	Balance3 float32
	Age2     int
	Age3     int8
	Age4     int16
	Age5     int32
	Age6     int32
	Age7     int64
	Age8     uint
	Age9     uint8
	Age10    uint16
	Age11    uint32
	Age12    uint64
	Date     time.Time
}

type ModelWithArray struct {
	Name    string
	Age     int64
	Address []*ModelForArray
}

type ModelForArray struct {
	Street string
	Number int16
}

func TestMarshallModelWithArray(t *testing.T) {

	add1 := &ModelForArray{Street:"Calle falsa", Number:123}
	add2 := &ModelForArray{Street:"Calle falsa2", Number:123}

	address := []*ModelForArray{add1,add2}

	mwa := ModelWithArray{
		Name:     "Fernando",
		Age:      32,
		Address: address,
	}

	m := ModelWithArray{}

	Marshall(mwa, &m)

	checkMarshallWithArray(mwa, m, t)
}

func TestMarshall(t *testing.T) {
	s := Sql{
		Name:     sql.NullString{String: "hola", Valid: true},
		Age:      sql.NullInt64{Int64: 1, Valid: true},
		Balance:  sql.NullFloat64{Float64: 1, Valid: true},
		Valid:    sql.NullBool{Bool: true, Valid: true},
		Name2:    "chau",
		Balance2: 1,
		Balance3: 1,
		Age2:     1,
		Age3:     1,
		Age4:     1,
		Age5:     1,
		Age6:     1,
		Age7:     1,
		Age8:     1,
		Age9:     1,
		Age10:    1,
		Age11:    1,
		Age12:    1,
		Date:     mysql.NullTime{Time: time.Now(), Valid: true},
	}
	m := Model{}
	Marshall(s, &m)

	checkMarshall(s, m, t)

}

func TestMarshallPointerOrigin(t *testing.T) {
	s := &Sql{
		Name:     sql.NullString{String: "hola", Valid: true},
		Age:      sql.NullInt64{Int64: 1, Valid: true},
		Balance:  sql.NullFloat64{Float64: 1, Valid: true},
		Valid:    sql.NullBool{Bool: true, Valid: true},
		Name2:    "chau",
		Balance2: 1,
		Balance3: 1,
		Age2:     1,
		Age3:     1,
		Age4:     1,
		Age5:     1,
		Age6:     1,
		Age7:     1,
		Age8:     1,
		Age9:     1,
		Age10:    1,
		Age11:    1,
		Age12:    1,
		Date:     mysql.NullTime{Time: time.Now(), Valid: true},
	}
	m := Model{}
	Marshall(s, &m)

	checkMarshall(*s, m, t)

}

func TestMarshallPointerDestino(t *testing.T) {
	s := Sql{
		Name:     sql.NullString{String: "hola", Valid: true},
		Age:      sql.NullInt64{Int64: 1, Valid: true},
		Balance:  sql.NullFloat64{Float64: 1, Valid: true},
		Valid:    sql.NullBool{Bool: true, Valid: true},
		Name2:    "chau",
		Balance2: 1,
		Balance3: 1,
		Age2:     1,
		Age3:     1,
		Age4:     1,
		Age5:     1,
		Age6:     1,
		Age7:     1,
		Age8:     1,
		Age9:     1,
		Age10:    1,
		Age11:    1,
		Age12:    1,
		Date:     mysql.NullTime{Time: time.Now(), Valid: true},
	}
	m := &Model{}
	Marshall(s, &m)

	checkMarshall(s, *m, t)

}

func TestMarshallSlice(t *testing.T) {
	s := Sql{
		Name:     sql.NullString{String: "hola", Valid: true},
		Age:      sql.NullInt64{Int64: 1, Valid: true},
		Balance:  sql.NullFloat64{Float64: 1, Valid: true},
		Valid:    sql.NullBool{Bool: true, Valid: true},
		Name2:    "chau",
		Balance2: 1,
		Balance3: 1,
		Age2:     1,
		Age3:     1,
		Age4:     1,
		Age5:     1,
		Age6:     1,
		Age7:     1,
		Age8:     1,
		Age9:     1,
		Age10:    1,
		Age11:    1,
		Age12:    1,
		Date:     mysql.NullTime{Time: time.Now(), Valid: true},
	}

	s1 := []Sql{}
	m1 := []Model{}
	s1 = append(s1, s)
	s1 = append(s1, s)
	s1 = append(s1, s)

	Marshall(s1, &m1)
	if len(s1) != len(m1) {
		t.Fatal("should be equals")
	}
	for index, elem := range s1 {
		checkMarshall(elem, m1[index], t)
	}

}

func TestMarshallPointerSliceOrigin(t *testing.T) {
	s := &Sql{
		Name:     sql.NullString{String: "hola", Valid: true},
		Age:      sql.NullInt64{Int64: 1, Valid: true},
		Balance:  sql.NullFloat64{Float64: 1, Valid: true},
		Valid:    sql.NullBool{Bool: true, Valid: true},
		Name2:    "chau",
		Balance2: 1,
		Balance3: 1,
		Age2:     1,
		Age3:     1,
		Age4:     1,
		Age5:     1,
		Age6:     1,
		Age7:     1,
		Age8:     1,
		Age9:     1,
		Age10:    1,
		Age11:    1,
		Age12:    1,
		Date:     mysql.NullTime{Time: time.Now(), Valid: true},
	}

	s1 := []*Sql{}
	m1 := []Model{}
	s1 = append(s1, s)
	s1 = append(s1, s)
	s1 = append(s1, s)

	Marshall(s1, &m1)
	if len(s1) != len(m1) {
		t.Fatal("should be equals")
	}
	for index, elem := range s1 {
		checkMarshall(*elem, m1[index], t)
	}

}

func TestMarshallPointerSliceDestino(t *testing.T) {
	s := Sql{
		Name:     sql.NullString{String: "hola", Valid: true},
		Age:      sql.NullInt64{Int64: 1, Valid: true},
		Balance:  sql.NullFloat64{Float64: 1, Valid: true},
		Valid:    sql.NullBool{Bool: true, Valid: true},
		Name2:    "chau",
		Balance2: 1,
		Balance3: 1,
		Age2:     1,
		Age3:     1,
		Age4:     1,
		Age5:     1,
		Age6:     1,
		Age7:     1,
		Age8:     1,
		Age9:     1,
		Age10:    1,
		Age11:    1,
		Age12:    1,
		Date:     mysql.NullTime{Time: time.Now(), Valid: true},
	}

	s1 := []Sql{}
	m1 := []*Model{}
	s1 = append(s1, s)
	s1 = append(s1, s)
	s1 = append(s1, s)

	Marshall(s1, &m1)
	if len(s1) != len(m1) {
		t.Fatal("should be equals")
	}
	for index, elem := range s1 {
		checkMarshall(elem, *m1[index], t)
	}
}

func TestMarshallSqlSlice(t *testing.T) {
	m := Model{"hola", 1, 1, true, "chau", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, time.Now()}
	s1 := []Sql{}
	m1 := []Model{}
	m1 = append(m1, m)
	m1 = append(m1, m)
	m1 = append(m1, m)
	Marshall(m1, &s1)
	if len(s1) != len(m1) {
		t.Fatal("should be equals")
	}
	for index, elem := range s1 {
		checkMarshallSql(elem, m1[index], t)
	}
}

func TestMarshallSqlPointerSliceOrigin(t *testing.T) {
	m := Model{"hola", 1, 1, true, "chau", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, time.Now()}
	s1 := []*Sql{}
	m1 := []Model{}
	m1 = append(m1, m)
	m1 = append(m1, m)
	m1 = append(m1, m)
	Marshall(m1, &s1)
	if len(s1) != len(m1) {
		t.Fatal("should be equals")
	}
	for index, elem := range s1 {
		checkMarshallSql(*elem, m1[index], t)
	}
}

func TestMarshallSqlPointerSliceDestino(t *testing.T) {
	m := &Model{"hola", 1, 1, true, "chau", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, time.Now()}
	s1 := []Sql{}
	m1 := []*Model{}
	m1 = append(m1, m)
	m1 = append(m1, m)
	m1 = append(m1, m)
	Marshall(m1, &s1)
	if len(s1) != len(m1) {
		t.Fatal("should be equals")
	}
	for index, elem := range s1 {
		checkMarshallSql(elem, *m1[index], t)
	}
}

func TestMarshallSql(t *testing.T) {
	s := Sql{}
	m := Model{"hola", 1, 1, true, "chau", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, time.Now()}
	Marshall(m, &s)
	checkMarshallSql(s, m, t)
}

func TestMarshallSqlPointerOrigin(t *testing.T) {
	s := &Sql{}
	m := Model{"hola", 1, 1, true, "chau", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, time.Now()}
	Marshall(m, &s)
	checkMarshallSql(*s, m, t)
}

func TestMarshallSqlPointerDestino(t *testing.T) {
	s := Sql{}
	m := &Model{"hola", 1, 1, true, "chau", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, time.Now()}
	Marshall(m, &s)
	checkMarshallSql(s, *m, t)
}

func checkMarshall(s Sql, m Model, t *testing.T) {
	if s.Date.Time != m.Date {
		t.Fatal("date should be equals")
	}
	value, _ := s.Name.Value()
	if value != m.Name {
		t.Fatal("value string should be equals")
	}
	value, _ = s.Age.Value()
	if value != m.Age {
		t.Fatal("value and string should be equals")
	}
	value, _ = s.Balance.Value()
	if value != m.Balance {
		t.Fatal("value and string should be equals")
	}
	value, _ = s.Valid.Value()
	if value != m.Valid {
		t.Fatal("value and string should be equals")
	}
	if s.Age2 != m.Age2 {
		t.Fatal("value and string should be equals")
	}
	if s.Age3 != m.Age3 {
		t.Fatal("value and string should be equals")
	}
	if s.Age4 != m.Age4 {
		t.Fatal("value and string should be equals")
	}
	if s.Age5 != m.Age5 {
		t.Fatal("value and string should be equals")
	}
	if s.Age6 != m.Age6 {
		t.Fatal("value and string should be equals")
	}
	if s.Age7 != m.Age7 {
		t.Fatal("value and string should be equals")
	}
	if s.Age8 != m.Age8 {
		t.Fatal("value and string should be equals")
	}
	if s.Age9 != m.Age9 {
		t.Fatal("value and string should be equals")
	}
	if s.Age10 != m.Age10 {
		t.Fatal("value and string should be equals")
	}
	if s.Age11 != m.Age11 {
		t.Fatal("value and string should be equals")
	}
	if s.Age12 != m.Age12 {
		t.Fatal("value and string should be equals")
	}
}

func checkMarshallWithArray(mwa ModelWithArray, m ModelWithArray, t *testing.T) {
	if mwa.Name != m.Name{
		t.Fatal("name should be equals")
	}
}

func checkMarshallSql(s Sql, m Model, t *testing.T) {
	value, _ := s.Name.Value()
	if value != m.Name {
		t.Fatal("value and string should be equals")
	}
	value, _ = s.Age.Value()
	if value != m.Age {
		t.Fatal("value and string should be equals")
	}
	value, _ = s.Balance.Value()
	if value != m.Balance {
		t.Fatal("value and string should be equals")
	}
	value, _ = s.Valid.Value()
	if value != m.Valid {
		t.Fatal("value and string should be equals")
	}
	if s.Age2 != m.Age2 {
		t.Fatal("value and string should be equals")
	}
	if s.Age3 != m.Age3 {
		t.Fatal("value and string should be equals")
	}
	if s.Age4 != m.Age4 {
		t.Fatal("value and string should be equals")
	}
	if s.Age5 != m.Age5 {
		t.Fatal("value and string should be equals")
	}
	if s.Age6 != m.Age6 {
		t.Fatal("value and string should be equals")
	}
	if s.Age7 != m.Age7 {
		t.Fatal("value and string should be equals")
	}
	if s.Age8 != m.Age8 {
		t.Fatal("value and string should be equals")
	}
	if s.Age9 != m.Age9 {
		t.Fatal("value and string should be equals")
	}
	if s.Age10 != m.Age10 {
		t.Fatal("value and string should be equals")
	}
	if s.Age11 != m.Age11 {
		t.Fatal("value and string should be equals")
	}
	if s.Age12 != m.Age12 {
		t.Fatal("value and string should be equals")
	}
}
