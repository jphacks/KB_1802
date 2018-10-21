package notify

import (
	"fmt"
	"github.com/KB_1802/notifyServer/constant"
	"github.com/line/line-bot-sdk-go/linebot"
	"math/rand"
)

var botText = []string{
	"(*ﾟーﾟ) 　お兄ちゃん\n実はこっそりお部屋に遊びに行きました。\nお兄ちゃんらしい部屋だと思うけど、ちゃんと片付けしないとダメだよ？",
	"J( 'ー`)し　たかしへ\n最近部屋の掃除をしてないんじゃないかい？\n忙しいかもしれないけど、早めに掃除するんだよ"}

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
