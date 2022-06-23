package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	DB *sql.DB
}

var Conn = &Connection{}

func (c *Connection) Connect() {
	var err error
	c.DB, err = sql.Open("mysql", "root:root@/wallet")
	if err != nil {
		panic(err)
	}
	c.DB.SetConnMaxLifetime(time.Minute * 3)
	c.DB.SetMaxOpenConns(10)
	c.DB.SetMaxIdleConns(10)
}
