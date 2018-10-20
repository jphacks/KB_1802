package api

import (
	"encoding/json"
	"fmt"
	"github.com/KB_1802/notifyServer/constant"
	"github.com/KB_1802/notifyServer/notify"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
LINEにログインする処理
内部的にはLINEのoauthにリダイレクトするだけ
 */
func Login(c *gin.Context) {

	oauthUrl := getOauthUrl()

	c.Redirect(302, oauthUrl)
	return
}

/*
LNEのoauth URLを取得
 */
func getOauthUrl() string {
	endpoint := "https://access.line.me/oauth2/v2.1/authorize"
	respType := "response_type=code"
	id := "client_id=" + constant.LoginId
	redUrl := "redirect_uri=" + constant.RedirectUrlEncoded
	state := "state=" + "12345abcde"
	scope := "scope=" + "openid%20profile"

	query := strings.Join([]string{respType, id, redUrl, state, scope}, "&")
	url := strings.Join([]string{endpoint, query}, "?")

	fmt.Println(url)

	return url

}


/*
ログインユーザーの内部IDを取得する
LINEのoauth認証からリダイレクトで返ってくる想定
 */
func GetUserId(c *gin.Context) {

	code := c.Query("code")

	if code == "" {
		fmt.Println()
	}

	//body := tokenRequest{grantType, code, redirectUrl, loginId, loginSecret}
	//fmt.Println(body)

	form := url.Values{}
	form.Add("grant_type", constant.GrantType)
	form.Add("code", code)
	form.Add("redirect_uri", constant.RedirectUrl)
	form.Add("client_id", constant.LoginId)
	form.Add("client_secret", constant.LoginSecret)

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
	//_, err = jwt.ParseWithClaims(idToken, claims, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(loginId), nil
	//})
	_, err = jwt.ParseWithClaims(idToken, claims, nil)
	// ... error handling
	if err != nil {
		fmt.Println(err)
	}

	userId := claims["sub"].(string)
	//fmt.Println(userId)

	notify.SendLINEMessage(userId)
}

/*
画像返却テスト
 */
func GetImage(c *gin.Context) {

	c.Status(http.StatusOK)
	c.File(constant.ImagePath + "/header20181020-144645.jpg")

}