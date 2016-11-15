package main

func main() {
	println("\nSqlite3:")
	sqliteDemo()

	println("\nPostgres with JsonB:")
	postgresJsonBDemo()

	println("\nMysql:")
	mysqlDemo()

	println("\nArangoDB:")
	aragnoDemo()

	println("\nElasticsearch:")
	elasticsearchDemo()
}
