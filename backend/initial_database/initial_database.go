package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "588800"
	dbname   = "task"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	sqlDropTable := `DROP TABLE IF EXISTS participant`
	_, err = db.Exec(sqlDropTable)
	if err != nil {
		panic(err)
	}

	sqlTable := `CREATE TABLE participant (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE NOT NULL,
		slot TEXT
		);`

	_, err = db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}
	fmt.Println("Added Participant Table successfully")

	sqlDropTable2 := `DROP TABLE IF EXISTS meeting`
	_, err = db.Exec(sqlDropTable2)
	if err != nil {
		panic(err)
	}

	sqlTable2 := `CREATE TABLE meeting (
		pid SERIAL PRIMARY KEY,
		participant1 INT NOT NULL,
		participant2 INT NOT NULL,
		startTime TIME,
		endTime TIME 
		);`

	_, err = db.Exec(sqlTable2)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Added Meeting Table successfully")
}
