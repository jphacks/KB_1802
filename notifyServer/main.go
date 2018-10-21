package main

import (
	"github.com/KB_1802/notifyServer/api"
	"github.com/KB_1802/notifyServer/constant"
	"github.com/KB_1802/notifyServer/db"
	"github.com/KB_1802/notifyServer/deployMode"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
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

	deployMode.Set()
	db.InitDB()
	rand.Seed(time.Now().UnixNano())

	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/twitter", api.TestTwitter)
	router.POST("/line", api.TestLine)

	router.POST("/dirtinessCheck", api.DirtinessCheck)


	router.GET("/login", api.Login)
	router.GET("/callback", api.GetUserId)
	router.GET("/image", api.GetImage)

	router.Run(":" + constant.Port)
}
