package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/KB_1802/notifyServer/db"
	"net/http"
)

func SendTwitterMessage() {

	//DBからレコード取得
	info := db.GetNotifyRecord()

	//json作成
	body := make(map[string]string)
	body["value1"] = info.TwitterMessage
	body["value2"] = "http://ec2-52-197-145-249.ap-northeast-1.compute.amazonaws.com:9000/image"

	//json marshal
	j, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//POST
	//contentType は text/json ではダメっぽい
	http.Post(info.TwitterUrl, "application/json", bytes.NewBuffer(j))

}
