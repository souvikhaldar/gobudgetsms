package gobudgetsms_test

import (
	"strings"
	"testing"

	"github.com/souvikhaldar/gobudgetsms"
)

func TestSendSMS(t *testing.T) {
	conf := gobudgetsms.SetConfig("", "", "", "", 0, 0, 0)
	res, er := gobudgetsms.SendSMS(conf, "", "", "")
	if er != nil {
		t.Error("Error in sending sms ", er)
	}
	fields := strings.Fields(res)
	if fields[0] != "OK" {
		t.Errorf("Error code %s, see https://www.budgetsms.net/sms-http-api/error-code/ for explanation of this error code", fields[1])
	}

}
