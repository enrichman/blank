package main

import (
	"database/sql"
	"fmt"
	"log"

	//
	// postgres
	_ "github.com/lib/pq"
	//
	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
	fmt.Println(version)
}

func PostgresConnectionString() string {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "password"
		dbname   = "postgres"
	)

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
}

func MysqlConnectionString() string {
	const (
		host     = "localhost"
		port     = 3306
		user     = "root"
		password = "password"
		dbname   = "mysql"
	)

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		user, password, host, port, dbname,
	)
}
