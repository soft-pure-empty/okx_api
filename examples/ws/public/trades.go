package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/ws/public"
)

func main() {
	handler := func(c public.EventTrades) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeTrades("BTC-USDT", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
