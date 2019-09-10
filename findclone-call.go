package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func findcloneCall() {
	req := newreq()
	req.SetRequestURI("https://findclone.ru/register?phone=+" + num)

	timeout := time.Minute + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `"pin_size"`) != -1 {
				okLog(callLog)
			} else {
				errLog(callLog)
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
