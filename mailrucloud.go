package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func mailruCloud() {
	req := newreq()
	req.SetBodyString(`{"phone":"+` + num + `","api":2,"email":"email","x-email":"x-email"}`)
	req.Header.SetContentType("application/json")
	req.SetRequestURI("https://cloud.mail.ru/api/v2/notify/applink")
	req.Header.SetMethod("POST")

	timeout := time.Second * 30
	var i uint8

	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for i < 3 && do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"status": 200`) != -1 {
				okLog(smsLog)
			} else {
				errLog(smsLog)
			}
			time.Sleep(timeout)
			i++
		}
	} else {
		for i < 3 && do {
			client.Do(req, nil)
			time.Sleep(timeout)
			i++
		}
	}
}
