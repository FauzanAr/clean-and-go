migrateup: 
	migrate -path db/migration -database "mysql://fauzan:passwordlocal123@tcp(localhost:3306)/watchDB" -verbose up

migratedown: 
	migrate -path db/migration -database "mysql://fauzan:passwordlocal123@tcp(localhost:3306)/watchDB" -verbose down