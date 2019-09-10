// +build !withoutTor

package main

import (
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cretz/bine/process"
	"github.com/fatih/color"
	proxy "github.com/valyala/fasthttp/fasthttpproxy"

	"github.com/ipsn/go-libtor"
)

func attachTor() {
	print("Запуск ")
	color.New(color.FgGreen).Print("Tor")
	println("...")

	var (
		err error

		i, status int

		n, port  string
		testsite = "http://example.com/"

		logOpt = "--quiet"

		tor process.Process

		torpath = filepath.Join(execDir, "tor")
		// torpath, _     = filepath.Abs("./tor-files-for-tests")
		torDataPath    = filepath.Join(torpath, "Data")
		torDataTorPath = filepath.Join(torDataPath, "Tor")

		geoipPath  = filepath.Join(torDataTorPath, "geoip")
		geoip6Path = filepath.Join(torDataTorPath, "geoip6")

		torDataDir string

		torrcPath = filepath.Join(torpath, "torrc")

		ln net.Listener
	)

	if os.Getenv("INFIBOMBTEST") == "1" {
		logOpt = "--hush"
	}

	for {
		if i == 9 {
			println("Не удалось запустить Tor! Попробуйте перезапустить программу!")
			shutdown(true)
		}

		n = strconv.Itoa(i)
		port = "376" + n

		ln, err = net.Listen("tcp", ":"+port)
		if err != nil || ln == nil {
			i++
			continue
		}
		ln.Close()

		torDataDir = filepath.Join(torDataPath, "data"+n)

		tor, err = libtor.Creator.New(nil, "-f", torrcPath, "--DataDirectory", torDataDir, "--SOCKSPort", port, "--GeoIPFile", geoipPath, "--GeoIPv6File", geoip6Path, logOpt)
		errCheck(err)

		err = tor.Start()
		if err != nil {
			print("Не удалось запустить ")
			color.New(color.FgGreen).Print("Tor")
			println("!")

			errCheck(err)
		}

		// Прописывание Tor-proxy в http-клиенте
		client.Dial = proxy.FasthttpSocksDialer("127.0.0.1:" + port)

		for b := 0; b < 3; b++ {
			status, _, err = client.Get(nil, testsite)
			if status == 200 && err == nil {
				break
			}
			time.Sleep(time.Second * 5)
		}
		status, _, err = client.Get(nil, testsite)
		if status == 200 && err == nil {
			break
		} else {
			print("Не удалось запустить ")
			color.New(color.FgGreen).Print("Tor")
			println("!")

			if err != nil {
				errCheck(err)
			}
			shutdown(true)
		}
	}

	color.New(color.FgGreen).Print("Tor")
	println(" запущен! Вы в безопасности ;)")
}
