package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
)

var maxId int
var r = rand.New(rand.NewSource(42))

func openDb(dbtype, connection string) *sql.DB {
	db, err := sql.Open(dbtype, connection)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}

func insertCities(db *sql.DB, cmd string) {
	m := StartMeasure("insert cities")
	tx, err := db.Begin()
	ch(err)
	stmt := chs(tx.Prepare(cmd))
	maxId = IterateCities(func(city *City) {
		b, err := json.Marshal(city)
		ch(err)
		chr(stmt.Exec(city.Id, city.Id, string(b)))
		m.Action()
	})
	stmt.Close()
	tx.Commit()
	m.End()
}

func queryByRandomId(stmt *sql.Stmt) {
	queryId := int(r.Int31n(int32(maxId)))
	rows, err := stmt.Query(queryId)
	ch(err)
	if !rows.Next() {
		panic("no rows in result set")
	}
	var data string
	ch(rows.Scan(&data))
	city := &City{}
	ch(json.Unmarshal([]byte(data), city))
	if city.Id != queryId {
		panic(fmt.Sprintf("received wrong id, expected %v, got %v", city.Id, queryId))
	}
	rows.Close()
}

func execN(title string, count int, db *sql.DB, statement string, callback func(stmt *sql.Stmt)) {
	m := StartMeasure(title)
	stmt := chs(db.Prepare(statement))
	for i := 0; i < count; i++ {
		callback(stmt)
		m.Action()
	}
	stmt.Close()
	m.End()
}

func chr(res sql.Result, err error) sql.Result {
	if err != nil {
		panic(err)
	}
	return res
}

func chs(stmt *sql.Stmt, err error) *sql.Stmt {
	if err != nil {
		panic(err)
	}
	return stmt
}

func ch(err error) {
	if err != nil {
		panic(err)
	}
}
