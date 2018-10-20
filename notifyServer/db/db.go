package db

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	dbmapNotify  *gorp.DbMap
	dbmapCamData *gorp.DbMap
)

type NotifyInfo struct {
	// db tag lets you specify the column name if it differs from the struct field
	LINEId         string `db:"line_id,size:50"`
	TwitterUrl     string `db:"twitter_url,size:1024"`
	TwitterMessage string `db:"twitter_message,size:1024"`
	Created        int64
}

type CamData struct {
	// db tag lets you specify the column name if it differs from the struct field
	FilePath    string  `db:"FILEPATH"`
	CaptureDate []uint8 `db:"CAPTUREDATE"`
	IsClean     int64   `db:"isClean"`
	Dirtiness   []uint8   `db:"DIRTINESS"`
}

func InitDB() {

	initNotify()
	initCamData()

}

func initNotify() {
	//DBコネクション
	db, err := sql.Open("mysql", "soisy:boaboa@tcp(52.197.145.249:3306)/nagatotest")
	if err != nil {
		panic(err)
	}

	//MySQL用にinit
	dbmapNotify = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8"}}

	//テーブル情報を登録
	dbmapNotify.AddTableWithName(NotifyInfo{}, "notify")

	//テーブル作成
	err = dbmapNotify.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func initCamData() {
	//DBコネクション
	db, err := sql.Open("mysql", "soisy:boaboa@tcp(52.197.145.249:3306)/soisy")
	if err != nil {
		panic(err)
	}

	//MySQL用にinit
	dbmapCamData = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8"}}

	//テーブル情報を登録
	dbmapCamData.AddTableWithName(CamData{}, "camData")
}

func GetNotifyRecord() *NotifyInfo {

	var info NotifyInfo
	err := dbmapNotify.SelectOne(&info, "select * from notify")
	if err != nil {
		log.Fatalln(err)
	}

	return &info
}

func GetNewestCamDataRecord() *CamData {

	var cam CamData
	err := dbmapCamData.SelectOne(&cam, "SELECT * FROM camData ORDER BY CAPTUREDATE DESC LIMIT 1")
	if err != nil {
		log.Fatalln(err)
	}

	return &cam
}
