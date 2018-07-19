package gobudgetsms

import (
	"io/ioutil"
	"log"
	"net/http"
)

// SendSMS returns the response after requesting Budget SMS api and the error if any
func SendSMS(details Details, message, to, from string) (string, error) {
	url := buildURL(detail, message, to, from)
	log.Println("The URL formed is ", url)
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
