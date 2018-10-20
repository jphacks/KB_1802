package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nagato/line-test/api"
)

/*
処理の流れ
ユーザーのアクセストークンを発行するために、ユーザーにoauthを介してログインしてもらう
アクセストークンの情報を元にユーザーIDを取得し、必要に応じて通知を送れるようにする
一度ユーザーIDを取得すれば、

BOTからの通知は友達追加しないと来ない
ウェブページにQRコードを表示して、ユーザーに友達登録させる必要があるっぽい

BOTを「J( 'ー`)し」にして部屋の汚れ具合についてコメントするようにする
 */

func main() {
	//port := os.Getenv("PORT")
	port := "9000"

	router := gin.New()
	router.Use(gin.Logger())

	//この処理を追記
	//router.POST("/linetest", api.PostTest)
	//router.POST("/line", api.Post)

	router.GET("/login", api.Login)
	router.GET("/callback", api.GetAccessToken)

	router.Run(":" + port)
}

