package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// import (
// 	"database/sql"
// 	"fmt"
// 	"hexagonal-architecture-golang/src/bootstrap"
// )

// type Database struct {
// 	Conn *sql.DB
// }

// var singleton *Database

// func NewDatabaseConnection() (bootstrap.Bootstrap, error) {
// 	return &Database{}, nil
// }

// func (d *Database) Initialice() error {
// 	conn, err := sql.Open("mysql", "root:firevivaldixdzzz@tcp(127.0.0.1:3306)/testiando")
// 	if err != nil {
// 		return err
// 	}
// 	d.Conn = conn
// 	fmt.Println("Database connection successfully")
// 	return nil
// }

// func GetMysqlConnection() *Database {
// 	return singleton
// }

type MysqlConnection struct {
	Conn *sql.DB
}

var singleton *MysqlConnection

func GetMysqlConnection() *MysqlConnection {
	if singleton == nil {
		conn, err := sql.Open("mysql", "root:firevivaldixdzzz@tcp(127.0.0.1:3306)/testiando")
		if err != nil {
			panic(err)
		}
		singleton = &MysqlConnection{Conn: conn}
		log.Println("Mysql connection created")
	}
	return singleton
}
