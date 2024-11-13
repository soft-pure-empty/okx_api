package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples"
	"github.com/soft-pure-empty/okx_api/ws"
	"github.com/soft-pure-empty/okx_api/ws/private"
)

func main() {
	args := &ws.Args{
		InstType: "SPOT",
	}
	handler := func(c private.EventOrders) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := private.SubscribeOrders(args, examples.Auth, handler, handlerError); err != nil {
		panic(err)
	}
	select {}
}
