package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/KB_1802/notifyServer/constant"
	"net/http"
)

func SendTwitterMessage() {

	//json作成
	body := make(map[string]string)
	body["value1"] = "test twitter."

	//json marshal
	j, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//POST
	//contentType は text/json ではダメっぽい
	http.Post(constant.TwitterNotifyUrl, "application/json", bytes.NewBuffer(j))

}
