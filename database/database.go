package database

import (
	"fmt"

	// import postgresql drivers
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// const connectionString = "host=%s port=%d dbname=%s user=%s password=%s sslmode=disable connect_timeout=%d"
// "mysql", "root:Michel42@tcp(127.0.0.1:3306)/"

const connectionString = "%s:%s@tcp(%s:%d)/%s"

//NewConnection to a database
func NewConnection(dbHost string, dbPort int, dbName string, dbUser string, dbPassword string) (apparteDb *sqlx.DB, err error) {
	cnx := fmt.Sprintf(connectionString, dbUser, dbPassword, dbHost, dbPort, dbName)
	apparteDb, err = sqlx.Connect("mysql", cnx)
	if err != nil {
		return nil, err
	}

	// Open doesn't open a connection. Validate DSN data:
	err = apparteDb.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connection to database established !")

	return
}
