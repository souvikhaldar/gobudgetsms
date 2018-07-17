package gobudgetsms

import (
	"io/ioutil"
	"net/http"
)

func sendSMS(detail Details, message string) (string, error) {
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
