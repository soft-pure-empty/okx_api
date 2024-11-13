package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/ws/public"
)

func main() {
	handler := func(c public.EventPriceLimit) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribePriceLimit("LTC-USD-190628", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
