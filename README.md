# Clean and Go
## Description
This is mini project that implement Clean Architecture concept and implement Dependecy Injection. This mini project also using minimalize library, like only use standart http library from Go, and using only Database Driver without using any ORM.
## Run and Go
1. Go to Makefile and change the database credential, also change database credetial on server.go
2. Run the migration using the following command
```bash
$ make migrateup
```
3. Run the application using the following command
```bash
$ go run server.go
```
## Finish and Go
Don't forget to install or prepare Mysql database on board!
<a name="Clean and go">Go Up</a>