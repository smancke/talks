# Json in SQL

The most SQL databases support a JSON datatype.

## Common Concepts

* The Json is stored more or less in a string type, which is validated on insert.
* Funtions for extraction and manipulation of the data, e.g. `json_extract(data, '$.Id')`.
* Syntactic sugar over the functions, e.g. `data->'Id'`.

## Postgres JsonB Type

* Postgtres has two types: Json and JsonB.
* JsonB stores the Json in a optimized binary form.
* Flexible index: GIN (Generalized Inverted Index)
* A very nice syntax.

## Demo
### Run the Demo
```shell
export GOPATH=`pwd`
go get --tags json1 example
bin/example
```

### Benchmarking results
```
Sqlite3:
insert cities: 73784/sec
select by column id: 34459/sec
select by column id_no_key: 169/sec
select by json json_extract()': 38/sec
create function index
select by json json_extract()': 15721/sec

Postgres with JsonB:
insert cities: 11101/sec
select by column id: 10244/sec
select by column id_no_key: 115/sec
select by json data->'Id': 27/sec
create GIN index
select by json data->'Id': 28/sec
select by json data @> '{"Id": $1}': 6445/sec
create function index
select by json data->'Id': 5963/sec

Mysql:
insert cities: 10254/sec
select by column id: 9790/sec
select by column id_no_key: 29/sec
select by data->'$.Id': 20/sec

ArangoDB:
select by json Id: 1655/sec

Elasticsearch:
insert cities: 17840/sec
select by json Id: 2386/sec
```

## References
### Postgres
* https://www.postgresql.org/docs/9.5/static/datatype-json.html
* https://www.postgresql.org/docs/9.5/static/functions-json.html

### SQlite
* https://sqlite.org/json1.html

### Mysql
* https://dev.mysql.com/doc/refman/5.7/en/json.html

## Disclaimer
License of Code: MIT, License of Doku: BY-SA 3.0

Test data taken from:

https://www.maxmind.com/de/free-world-cities-database

http://download.maxmind.com/download/geoip/database/LICENSE_WC.txt
