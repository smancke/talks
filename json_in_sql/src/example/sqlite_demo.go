package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func sqliteDemo() {
	db := openDb("sqlite3", "/tmp/sqlite_demo.db")
	defer db.Close()

	chr(db.Exec(`drop table if exists demo`))
	chr(db.Exec(`
	          create table demo
	          (
	            id integer,
	            id_no_key integer,
	            data text,
	            primary key(id)
	          );
	        `))

	insertCities(db, "insert into demo (id, id_no_key, data) values ($1, $2, json($3))")

	execN("select by column id", 10000, db,
		`select data from demo where id = $1`,
		queryByRandomId)

	execN("select by column id_no_key", 100, db,
		`select data from demo where id_no_key = $1`,
		queryByRandomId)

	execN("select by json json_extract()'", 100, db,
		`select data from demo where json_extract(data, '$.Id') = $1`,
		queryByRandomId)

	println("create function index")
	chr(db.Exec(`CREATE INDEX id ON demo(json_extract(data, '$.Id'))`))

	execN("select by json json_extract()'", 100, db,
		`select data from demo where json_extract(data, '$.Id') = $1`,
		queryByRandomId)
}
