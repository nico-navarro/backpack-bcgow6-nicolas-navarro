package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnection() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME")))
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
