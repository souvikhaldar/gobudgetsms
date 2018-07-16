package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	msg := createmsg("Your OTP for Kartbites is 5467")

	// fill the parameters according to https://www.budgetsms.net/sms-http-api/test-sms/
	// details regarding account can be found here https://www.budgetsms.net/controlpanel/api-details/
	detail := details{
		"kartbites",
		"15433",
		"4c072defb2bdf7f54633dc8e71fa94ff",
		msg,
		"kartbites",
		"+918653149116",
		"",
		0,
		0,
		0,
	}
	path := "/testsms"
	url := buildURL(detail, path)
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
