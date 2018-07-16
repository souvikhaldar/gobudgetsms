package gobudgetsms

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var detail details

func setConfig(message string) {
	url := buildURL(detail)
	res, er := http.Get(url)
	if er != nil {
		fmt.Println("Error in sending the sms ", er)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading from body ", err)
		return
	}
	fmt.Println("Body: ", string(body))
}
