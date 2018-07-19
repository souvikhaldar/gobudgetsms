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

var detail Details

func createmsg(msg string) string {
	message := []rune(msg)
	for i := 0; i < len(message); i++ {
		if string(message[i]) == " " {
			message[i] = 43
		}
	}
	return string(message)
}
func buildURL(det Details, message, to, from string) string {
	u := &url.URL{
		Scheme:   "https",
		Host:     "api.budgetsms.net",
		Path:     "/sendsms",
		RawQuery: fmt.Sprintf("username=%s&userid=%s&handle=%s&msg=%s&from=%s&to=%s&customid=%s&price=%d&mccmnc=%d&credit=%d", det.Username, det.Userid, det.Handle, createmsg(message), from, to, det.Customid, det.Price, det.Mccmnc, det.Credit),
	}
	return u.String()
}

// SetConfig allows to enter the credentials of Budget sms account
func SetConfig(Username, Userid, Handle, Customid string, Price, Mccmnc, Credit int) Details {
	detail := Details{
		Username,
		Userid,
		Handle,
		"",
		"",
		"",
		"",
		Price,
		Mccmnc,
		Credit,
	}
	return detail
}
