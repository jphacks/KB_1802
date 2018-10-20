package notify

import (
	"fmt"
	"github.com/KB_1802/notifyServer/constant"
	"github.com/line/line-bot-sdk-go/linebot"

)

func SendLINEMessage (id string) {

	//channelSecret: 文字通り channelToken: アクセストークン
	bot, err := linebot.New(constant.BotSecret, constant.AccessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	var messages []linebot.SendingMessage
	// append some message to messages
	text := linebot.NewTextMessage("linebot test.")
	messages = append(messages, text)
	_, err = bot.PushMessage(id, messages...).Do()
	if err != nil {
		// Do something when some bad happened
		fmt.Println(err)
		return
	}

}
