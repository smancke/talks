package main

import (
	_ "github.com/lib/pq"
)

func postgresJsonBDemo() {
	db := openDb("postgres", "postgres://postgres:mysecretpassword@localhost/postgres?sslmode=disable")
	defer db.Close()

	chr(db.Exec(`drop table if exists demo`))
	chr(db.Exec(`
	          create table demo
	          (
	            id integer,
	            id_no_key integer,
	            data jsonb,
	            primary key(id)
	          );
	        `))

	insertCities(db, "insert into demo (id, id_no_key, data) values ($1, $2, $3)")

	execN("select by column id", 10000, db,
		`select data from demo where id = $1`,
		queryByRandomId)

	execN("select by column id_no_key", 100, db,
		`select data from demo where id_no_key = $1`,
		queryByRandomId)

	execN("select by json data->'Id'", 100, db,
		`select data from demo where data->'Id' = $1`,
		queryByRandomId)

	println("create GIN index")
	chr(db.Exec(`CREATE INDEX on demo USING GIN (data);`))

	execN("select by json data->'Id'", 100, db,
		`select data from demo where data->'Id' = $1`,
		queryByRandomId)

	execN(`select by json data @> '{"Id": $1}'`, 10000, db,
		`select data from demo where data @> json_build_object('Id', $1::integer)::jsonb`,
		queryByRandomId)

	println("create function index")
	chr(db.Exec(`CREATE INDEX id ON demo((data->'Id'))`))

	execN("select by json data->'Id'", 100, db,
		`select data from demo where data->'Id' = $1`,
		queryByRandomId)
}
