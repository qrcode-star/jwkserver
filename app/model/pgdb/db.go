package pgdb

import (
	"github.com/gogf/gf/database/gdb"
	_ "github.com/lib/pq"
)

func Newdb() (db gdb.DB, err error) {
	mydb := gdb.ConfigNode{}
	mydb.Type="pgsql"
	mydb.LinkInfo = "postgres://pg@172.18.100.57/pg?sslmode=disable"
	gdb.AddDefaultConfigNode(mydb)
	db ,err = gdb.New()
	return db,err
}