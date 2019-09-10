package main

import (
	"bufio"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

const version = "1.7.0"

func shutdown(crash bool) {
	print("Нажмите 'Enter', для закрытия программы...")
	bufio.NewReader(os.Stdin).ReadRune()

	if crash {
		os.Exit(1)
	}
	os.Exit(0)
}

func errCheck(err error) {
	if err != nil {
		println(err.Error())
		debug.PrintStack()

		shutdown(true)
	}
}

var client = &fasthttp.Client{
	DisableHeaderNamesNormalizing: true,
	NoDefaultUserAgentHeader:      true,
}

var (
	num       string
	floodMode int8
	logging   int8
	floodTime int

	do = true

	smss int
	mux  = sync.Mutex{}
)

type logParam bool

const (
	callLog logParam = true
	smsLog  logParam = false
)

var grPrntln = color.New(color.FgHiGreen).Println
var okLog = func(call logParam) {
	mux.Lock()
	if call {
		grPrntln(time.Now().Format("15:04:05.000") + " - Запрос на звонок отправлен!")
	} else {
		smss++
		grPrntln(time.Now().Format("15:04:05.000") + " - SMS отправлено! (" + strconv.Itoa(smss) + ")")
	}
	mux.Unlock()
}

var redPrntln = color.New(color.FgRed).Println
var errLog = func(call logParam) {
	mux.Lock()
	if call {
		redPrntln(time.Now().Format("15:04:05.000") + " - Не удалось отправить запрос на звонок!")
	} else {
		redPrntln(time.Now().Format("15:04:05.000") + " - SMS не отправлено!")
	}
	mux.Unlock()
}

func main() {
	color.New(color.FgHiGreen).Print("Запущен флуд на номер")
	print(" ")
	color.New(color.FgHiYellow).Println("+" + num)

	switch logging {
	case 0:
		println("Лог выключен")
	case 1:
		println("Лог будет содержать только OK сообщения")
	case 2:
		println("Лог будет содержать только ERR сообщения")
	case 3:
		println("Лог будет содержать OK и ERR сообщения")
	}

	switch floodMode {
	case 1:
		println("Будут отправляться запросы на отправку SMS")
		runSMSServices()
	case 2:
		println("Будут отправляться запросы на звонки")
		runCallServices()
	case 3:
		println("Будут отправляться запросы на отправку SMS и на звонки")
		runSMSServices()
		runCallServices()
	}

	if os.Getenv("INFIBOMBTEST") == "1" {
		go test()
	}

	color.New(color.FgHiYellow).Println("К сожалению (по техническим причинам), отправки смс через сервис без КД (wer.ru) не будут учитываться в счётчике")

	println("Нажмите Ctrl+C чтобы остановить работу бомбера")

	if floodTime <= 0 {
		<-make(chan bool, 0)
	} else {
		<-time.NewTimer(time.Second * time.Duration(floodTime)).C
		do = false
		mux.Lock()
		println("Работа завершена!")
		shutdown(false)
	}
}

func runSMSServices() {
	go apteka366()
	go belkacar()
	go deliclub()
	go deliclublink()
	go dodopizza()
	go dostaevsky()
	go drugvokrug()
	go ennergiia()
	go fex()
	go funday()
	go gett()
	go gorzdrav()
	go grab()
	go icq()
	go jacofood()
	go karusel()
	go kfc()
	go lenta()
	go mailruCloud()
	go masterhost()
	go mtstv()
	go mvideo()
	go stoloto()
	go tinder()
	go tinkoff()
	go viasat()
	go xarakiri()
	go yandeda()
	go youla()
	go zaymigo()
	go zvuk()

	go greadios()

	go wer()
	go wer()
}

func runCallServices() {
	go findcloneCall()
}

func newreq() (req *fasthttp.Request) {
	req = &fasthttp.Request{}
	req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	return
}
