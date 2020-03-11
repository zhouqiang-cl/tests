package main

import (
	"fmt"
	"flag"
	"database/sql"
	"context"

	"github.com/zhouqiang-cl/tests/tests/largetxn"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var conn struct {
		user string
		host  string
		port    int
		db string
		dsn string
	}

	flag.StringVar(&conn.host, "H", "", "host addr")
	flag.IntVar(&conn.port, "P", 4000, "port of TiDB")
	flag.StringVar(&conn.db, "B", "test", "The test database")

	conn.dsn = fmt.Sprintf("root:@tcp(%s:%d)/test", conn.host, conn.port) 

	db, err := sql.Open("mysql", conn.dsn)
	if err != nil {
		fmt.Println("create mysql conn failed+%v", err)
	}
	con, err := db.Conn(context.Background())
	if err != nil {
		fmt.Println(" create conn error +%v", err)
	}

	err = largetxn.LargeTxn(5000, con)
	if err != nil {
		fmt.Println(" create conn error +%v", err)
	}
}
