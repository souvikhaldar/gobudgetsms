package gobudgetsms

import (
	"io/ioutil"
	"net/http"
)

// SendSMS returns the response after requesting Budget SMS api and the error if any
func SendSMS(detail Details, message string) (string, error) {
	url := buildURL(detail)
	res, er := http.Get(url)
	if er != nil {
		return "", er
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
