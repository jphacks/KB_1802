package constant

type Mode int

const (
	local Mode = iota
	aws
)

const (
	Port = "9000"

	//LINE
	LoginId     = "1615648209"
	LoginSecret = "8ef1a13cabd721c5432abbf153b854b5"
	GrantType   = "authorization_code"
	BotId       = "1615598312"
	BotSecret   = "dfb8715b34782df18d93885948e3b624"
	AccessToken = "B/G/g+pSwBvNbUNWZKwCekcupsJ4lxYB+tWVl6TuixzNR4UFxL/Rw9RuQHZroRIXBNbQFArMbb5XWRNXuiinXIUljumrBTp8a+ALyvjRTK0p4tXXDFms6FHKqciUtD3F87wiXIxN3G4mLdpfzeV/rAdB04t89/1O/w1cDnyilFU="

	//Twitter
	TwitterNotifyUrl = "https://maker.ifttt.com/trigger/send_twitter/with/key/fRkVln270KZD_sfpqwuBrVISEXQTMDjM9ntQJP_GYLg"

	//GetImage
	ImagePath = "/home/ubuntu/camData"

	)

var (
	RedirectUrl        string
	RedirectUrlEncoded string
)

func SetRedirectUrl(ip string) {
	RedirectUrl = "http://" + ip + ":" + Port + "/callback"
	RedirectUrlEncoded = "http%3A%2F%2F" + ip + ":" + Port + "%2Fcallback"
}
