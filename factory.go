package gobudgetsms

import (
	"fmt"
	"net/url"
)

// Details has all the credentials of the Budget SMS account as well as the details regarding the SMS
type Details struct {
	Username string
	Userid   string
	Handle   string
	Msg      string
	From     string
	To       string
	Customid string
	Price    int
	Mccmnc   int
	Credit   int
}

func createmsg(msg string) string {
	message := []rune(msg)
	for i := 0; i < len(message); i++ {
		if string(message[i]) == " " {
			message[i] = 43
		}
	}
	return string(message)
}
func buildURL(det Details) string {
	u := &url.URL{
		Scheme:   "https",
		Host:     "api.budgetsms.net",
		Path:     "/sendsms",
		RawQuery: fmt.Sprintf("username=%s&userid=%s&handle=%s&msg=%s&from=%s&to=%s&customid=%s&price=%d&mccmnc=%d&credit=%d", det.Username, det.Userid, det.Handle, createmsg(det.Msg), det.From, det.To, det.Customid, det.Price, det.Mccmnc, det.Credit),
	}
	return u.String()
}
