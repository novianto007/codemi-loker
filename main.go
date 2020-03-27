package main

import (
	"bufio"
	"codemi/loker/actions"
	"codemi/loker/handler"
	"codemi/loker/helpers"
	"os"
)

func init() {
	handler.RegisterHandler("init", actions.Init)
	handler.RegisterHandler("input", actions.Input)
	handler.RegisterHandler("status", actions.Status)
	handler.RegisterHandler("leave", actions.Leave)
	handler.RegisterHandler("find", actions.Find)
	handler.RegisterHandler("search", actions.Search)
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	helpers.Println("selamat data di diloker app")
	for s.Scan() {
		if s.Text() == "exit" {
			break
		}
		handler.Handle(s.Text())
	}
	helpers.Println("sampai jumpa")
}
