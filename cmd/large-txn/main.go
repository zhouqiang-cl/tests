package main

import (
	"fmt"
	"flag"
	"database/sql"

	"github.com/zhouqiang-cl/tests/tests/largetxn"

)

func main() {

	var conn struct {
		user string
		host  string
		port    int
		db string
		dsn string
	}

	flag.StringVar(&conn.host, "H", "127.0.0.1", "host addr")
	flag.IntVar(&conn.port, "P", 4000, "port of TiDB")
	flag.StringVar(&conn.db, "B", "test", "The test database")

	conn.dsn = fmt.Sprintf("root:@tcp(%s:%d)/test", conn.host, conn.port) 

	db, err := sql.Open("mysql", conn.dsn)
	if err != nil {
		fmt.Println("create mysql conn failed")
	}

	largetxn.LargeTxn(5000, db)
}