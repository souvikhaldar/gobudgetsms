package gobudgetsms

import (
	"fmt"
	"net/url"
)

// Details has all the credentials of the Budget SMS account as well as the details regarding the SMS
type Details struct {
	username string
	userid   string
	handle   string
	msg      string
	from     string
	to       string
	customid string
	price    int
	mccmnc   int
	credit   int
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
		RawQuery: fmt.Sprintf("username=%s&userid=%s&handle=%s&msg=%s&from=%s&to=%s&customid=%s&price=%d&mccmnc=%d&credit=%d", det.username, det.userid, det.handle, createmsg(det.msg), det.from, det.to, det.customid, det.price, det.mccmnc, det.credit),
	}
	fmt.Println("Final URL is", u.String())
	return u.String()
}
