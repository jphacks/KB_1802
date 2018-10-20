package deployMode

import (
	"flag"
	"github.com/KB_1802/notifyServer/constant"
)

func Set() {
	var mode string
	flag.StringVar(&mode, "deploy", "local", "set deploy mode")
	flag.Parse()

	ip := "127.0.0.1"
	if mode == "local" {
		ip = "127.0.0.1"
	} else if mode == "aws" {
		ip = "52.197.145.249"
	} else {
		panic("mode is illegal. mode should be [aws, local]")
	}

	constant.SetRedirectUrl(ip)

}
