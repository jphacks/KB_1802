package notify

import (
	"fmt"
	"github.com/KB_1802/notifyServer/constant"
	"github.com/line/line-bot-sdk-go/linebot"
	"math/rand"
)

var botText = []string{
	"J( 'ー`)し　たかしへ\nちゃんと掃除しているかい？",
	"J( 'ー`)し　たかしへ\n最近部屋の掃除をしてないんじゃないかい？"}

func SendLINEMessage (id string) {

	//channelSecret: 文字通り channelToken: アクセストークン
	bot, err := linebot.New(constant.BotSecret, constant.AccessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	var messages []linebot.SendingMessage
	// append some message to messages
	randomText := botText[rand.Intn(len(botText))]
	text := linebot.NewTextMessage(randomText)
	messages = append(messages, text)
	_, err = bot.PushMessage(id, messages...).Do()
	if err != nil {
		// Do something when some bad happened
		fmt.Println(err)
		return
	}

}
