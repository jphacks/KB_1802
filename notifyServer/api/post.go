package api

import (
	"fmt"
	"github.com/KB_1802/notifyServer/db"
	"github.com/KB_1802/notifyServer/notify"
	"github.com/gin-gonic/gin"
	"strconv"
)

type tokenRequest struct {
	grant    string `json:"grant_type"`
	code     string `json:"code"`
	redirect string `json:"redirect_uri"`
	id       string `json:"client_id"`
	secret   string `json:"client_secret"`
}



func TestTwitter(c *gin.Context) {

	notify.SendTwitterMessage()

}

func TestLine(c *gin.Context) {

	info := db.GetNotifyRecord()
	notify.SendLINEMessage(info.LINEId)
}

func DirtinessCheck(c *gin.Context) {

	level := c.Query("level")

	if level == "" {
		fmt.Println("level is empty.")
		return
	}

	dirtiness, err := strconv.Atoi(level)
	if err != nil {
		fmt.Println("cannot cast 'level' to int.")
		return
	}
	//image := nil //test



	//imageどうする？
	if dirtiness == 2 { //twitter
		notify.SendTwitterMessage()
	} else if dirtiness == 1 { //LINE
		info := db.GetNotifyRecord()
		notify.SendLINEMessage(info.LINEId)
	}

	fmt.Println("call dirtinessCheck")
	c.AbortWithStatus(204)

}

//func PostTest(c *gin.Context) {
//
//	//botSecret: 文字通り channelToken: アクセストークン
//	bot, err := linebot.New(botSecret, accessToken)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	var messages []linebot.SendingMessage
//	// append some message to messages
//	text := linebot.NewTextMessage("linebot test.")
//	messages = append(messages, text)
//	_, err = bot.PushMessage("Uba5808517e79d4e19d386798228b386f", messages...).Do()
//	if err != nil {
//		// Do something when some bad happened
//		fmt.Println(err)
//		return
//	}
//}



//func Post(c *gin.Context) {
//	//proxyURL, _ := url.Parse(os.Getenv("FIXIE_URL"))
//	//client := &http.Client{
//	//	Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
//	//}
//
//	//アクセストークン取得
//	accessToken, err := redirectLogin()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//
//}



//func redirectLogin() (string, error) {
//
//	oauthUrl := getOauthUrl()
//
//	//body := oauth{grantType, botId, botSecret}
//	//j, err := json.Marshal(body)
//	//if err != nil {
//	//	return "", err
//	//}
//	//buf := bytes.NewBuffer(j)
//
//	req, err := http.NewRequest(http.MethodGet, oauthUrl, nil)
//	if err != nil {
//		return "", err
//	}
//
//	c := &http.Client{}
//	resp, err := c.Do(req)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	respBody, _ := ioutil.ReadAll(resp.Body)
//	responseJson := make(map[string]interface{})
//	err = json.Unmarshal(respBody, &responseJson)
//	if err != nil {
//		return "", err
//	}
//
//	fmt.Println(responseJson)
//
//	accessToken := responseJson["access_token"].(string)
//	fmt.Println("accessToken : ", accessToken)
//
//
//
//
//	return accessToken, nil
//
//}
//
//
