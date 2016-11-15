package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func mysqlDemo() {
	db := openDb("mysql", "demo:demo@/demo")
	defer db.Close()

	chr(db.Exec(`drop table if exists demo`))
	chr(db.Exec(`
          create table demo
          (
            id int not null key,
            id_no_key int,
            data json
          );
        `))

	insertCities(db, "insert into demo (id, id_no_key, data) values (?, ?, ?)")

	execN("select by column id", 10000, db,
		`select data from demo where id = ?`,
		queryByRandomId)

	execN("select by column id_no_key", 100, db,
		`select data from demo where id_no_key = ?`,
		queryByRandomId)

	execN("select by data->'$.Id'", 100, db,
		`select data from demo where data->'$.Id' = ?`,
		queryByRandomId)
}
