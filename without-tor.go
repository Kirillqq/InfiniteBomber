// +build withoutTor

package main

import "github.com/fatih/color"

func attachTor() {
	color.New(color.FgGreen).Print("Tor")
	print(" не будет запущен! Вы НЕ в безопасности :[\nЧтобы ")
	color.New(color.FgGreen).Print("Tor")
	println(" использовался, необходимо использовать соответствующую версию программы!")
}
