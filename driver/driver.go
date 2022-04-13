package driver

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(dbConnectionString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
