package db

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dbmap *gorp.DbMap

type NotifyInfo struct {
	// db tag lets you specify the column name if it differs from the struct field
	LINEId         string `db:"line_id,size:50"`
	TwitterUrl     string `db:"twitter_url,size:1024"`
	TwitterMessage string `db:"twitter_message,size:1024"`
	Created        int64
}

func InitDB() {

	//DBコネクション
	db, err := sql.Open("mysql", "soisy:boaboa@tcp(52.197.145.249:3306)/nagatotest")
	if err != nil {
		panic(err)
	}

	//MySQL用にinit
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8"}}

	//テーブル情報を登録
	dbmap.AddTableWithName(NotifyInfo{}, "notify")

	//テーブル作成
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func GetNotifyRecord() *NotifyInfo {

	var info NotifyInfo
	err := dbmap.SelectOne(&info, "select * from notify")
	if err != nil {
		log.Fatalln(err)
	}

	return &info
}
