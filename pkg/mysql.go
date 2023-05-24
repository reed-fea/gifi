package pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func init_mysql() {
	// fmt.Println("init pkg")
	connStr := "user=sgmnqtui password=H94izUPgL0M8ue111kMF3gvJkeeIWbIN dbname=sgmnqtui host=tiny.db.elephantsql.com port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// age := 27
	// rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	// create a table, fields: id, name, age
	// rows, err := db.Query("CREATE TABLE users (id serial PRIMARY KEY, name VARCHAR (50) NOT NULL, age INT NOT NULL)")
	// insert a row into table random name and age
	rows, err := db.Query("INSERT INTO users (name, age) VALUES ('random', 27)")
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(rows)
}

func Run() {
	init_mysql()
}
