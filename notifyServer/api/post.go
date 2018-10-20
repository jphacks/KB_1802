package api

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)
type Mode int

const (
	local Mode = iota
	aws
)


var (

	mode = aws

	loginId            = "1615648209"
	loginSecret        = "8ef1a13cabd721c5432abbf153b854b5"
	grantType          = "authorization_code"
	botId              = "1615598312"
	botSecret          = "dfb8715b34782df18d93885948e3b624"
	redirectUrl        = "http://127.0.0.1:9000/callback"
	redirectUrlEncoded = "http%3A%2F%2F127.0.0.1:9000%2Fcallback"
	accessToken        = "B/G/g+pSwBvNbUNWZKwCekcupsJ4lxYB+tWVl6TuixzNR4UFxL/Rw9RuQHZroRIXBNbQFArMbb5XWRNXuiinXIUljumrBTp8a+ALyvjRTK0p4tXXDFms6FHKqciUtD3F87wiXIxN3G4mLdpfzeV/rAdB04t89/1O/w1cDnyilFU="
)

type tokenRequest struct {
	grant    string `json:"grant_type"`
	code     string `json:"code"`
	redirect string `json:"redirect_uri"`
	id       string `json:"client_id"`
	secret   string `json:"client_secret"`
}

func init() {
	if mode == aws {
		redirectUrl        = "http://52.197.145.249:9000/callback"
		redirectUrlEncoded = "http%3A%2F%2F52.197.145.249:9000%2Fcallback"
	}
}

func PostTest(c *gin.Context) {

	//botSecret: 文字通り channelToken: アクセストークン
	bot, err := linebot.New(botSecret, accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	var messages []linebot.SendingMessage
	// append some message to messages
	text := linebot.NewTextMessage("linebot test.")
	messages = append(messages, text)
	_, err = bot.PushMessage("Uba5808517e79d4e19d386798228b386f", messages...).Do()
	if err != nil {
		// Do something when some bad happened
		fmt.Println(err)
		return
	}

}

func Post(c *gin.Context) {
	//proxyURL, _ := url.Parse(os.Getenv("FIXIE_URL"))
	//client := &http.Client{
	//	Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
	//}

	//アクセストークン取得
	accessToken, err := redirectLogin()
	if err != nil {
		fmt.Println(err)
		return
	}

	//botSecret: 文字通り channelToken: アクセストークン
	bot, err := linebot.New(botSecret, accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	var messages []linebot.SendingMessage
	// append some message to messages
	text := linebot.NewTextMessage("linebot test.")
	messages = append(messages, text)
	_, err = bot.PushMessage("Uba5808517e79d4e19d386798228b386f", messages...).Do()
	if err != nil {
		// Do something when some bad happened
		fmt.Println(err)
		return
	}
}

func GetAccessToken(c *gin.Context) {

	code := c.Query("code")

	//body := tokenRequest{grantType, code, redirectUrl, loginId, loginSecret}
	//fmt.Println(body)

	form := url.Values{}
	form.Add("grant_type",grantType)
	form.Add("code",code)
	form.Add("redirect_uri",redirectUrl)
	form.Add("client_id",loginId)
	form.Add("client_secret",loginSecret)

	body := strings.NewReader(form.Encode())

	resp, err := http.Post("https://api.line.me/oauth2/v2.1/token", "application/x-www-form-urlencoded", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()




	respBody, _ := ioutil.ReadAll(resp.Body)
	responseJson := make(map[string]interface{})
	err = json.Unmarshal(respBody, &responseJson)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode >= 400 {
		var foo interface{}
		err = json.NewDecoder(resp.Body).Decode(&foo)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(foo)
		return
	}

	idToken := responseJson["id_token"].(string)
	//fmt.Println(idToken)

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(idToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(loginId), nil
	})
	// ... error handling
	if err != nil {
		fmt.Println(err)
	}

	userId := claims["sub"].(string)
	//fmt.Println(userId)

	//botSecret: 文字通り channelToken: アクセストークン
	bot, err := linebot.New(botSecret, accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	var messages []linebot.SendingMessage
	// append some message to messages
	text := linebot.NewTextMessage("linebot test.")
	messages = append(messages, text)
	_, err = bot.PushMessage(userId, messages...).Do()
	if err != nil {
		// Do something when some bad happened
		fmt.Println(err)
		return
	}



}

func Login(c *gin.Context) {

	oauthUrl := getOauthUrl()

	c.Redirect(302, oauthUrl)
	return
}

func redirectLogin() (string, error) {

	oauthUrl := getOauthUrl()

	//body := oauth{grantType, botId, botSecret}
	//j, err := json.Marshal(body)
	//if err != nil {
	//	return "", err
	//}
	//buf := bytes.NewBuffer(j)

	req, err := http.NewRequest(http.MethodGet, oauthUrl, nil)
	if err != nil {
		return "", err
	}

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	responseJson := make(map[string]interface{})
	err = json.Unmarshal(respBody, &responseJson)
	if err != nil {
		return "", err
	}

	fmt.Println(responseJson)

	accessToken := responseJson["access_token"].(string)
	fmt.Println("accessToken : ", accessToken)




	return accessToken, nil

}

func getOauthUrl() string {
	//https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=1234567890&redirect_uri=https%3A%2F%2Fexample.com%2Fauth&state=12345abcde&scope=openid%20profile&nonce=09876xyz
	endpoint := "https://access.line.me/oauth2/v2.1/authorize"
	respType := "response_type=code"
	id := "client_id=" + loginId
	redUrl := "redirect_uri=" + redirectUrlEncoded
	state := "state=" + "12345abcde"
	scope := "scope=" + "openid%20profile"

	query := strings.Join([]string{respType, id, redUrl, state, scope}, "&")
	url := strings.Join([]string{endpoint, query}, "?")

	fmt.Println(url)

	return url

}


