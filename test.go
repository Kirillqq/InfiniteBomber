package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func test() {
	req := newreq()
	req.SetBodyString(`{"phone_number":"` + num + `"}`)
	req.Header.SetContentType("application/json")
	req.SetRequestURI("https://httpbin.org/post")
	req.Header.SetMethod("POST")

	timeout := time.Second * 12
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			println(req.URI().String(), err.Error(), string(res.Body()))

			if err == nil && strings.Index(string(res.Body()), `origin`) != -1 {
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
