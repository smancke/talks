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
