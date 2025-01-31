package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples"
	"github.com/soft-pure-empty/okx_api/ws/private"
)

func main() {
	handler := func(c private.EventBalanceAndPosition) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := private.SubscribeBalanceAndPosition(examples.Auth, handler, handlerError); err != nil {
		panic(err)
	}
	select {}
}
