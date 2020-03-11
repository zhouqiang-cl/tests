package largetxn

import (
	"fmt"
	"database/sql"
	"math/rand"
	"context"
)

// LargeTxn is for large transaction
func LargeTxn(cnt int, db *sql.DB) error {
	fmt.Println("run large txn DML")

	conn, err := db.Conn(context.Background())

	tx, err := conn.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	for i := 0; i < cnt; i++ {
		_, err := tx.Exec("update sbtest1 set k = k + 1 where id in (?, ?, ?)", rand.Intn(100000), rand.Intn(100000), rand.Intn(100000))
		if err != nil {
			return err
		}
		if i%1000 == 0 {
			fmt.Println("current executed = ", i)
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("run large txn DML finish")
	return nil
}