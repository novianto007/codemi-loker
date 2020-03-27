package handler

import (
	"codemi/loker/helpers"
	"strings"
)

var routeHandler map[string]func([]string) = map[string]func([]string){}

func RegisterHandler(cmd string, act func([]string)) {
	routeHandler[cmd] = act
}

func Handle(text string) {
	cmd, params := filterInput(text)
	if act, ok := routeHandler[strings.ToLower(cmd)]; ok {
		act(params)
	} else {
		helpers.Println("perintah tidak valid")
	}
}

func filterInput(text string) (string, []string) {
	tmp := strings.Split(text, " ")
	return tmp[0], tmp[1:]
}
