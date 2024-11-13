package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/ws"
	"github.com/soft-pure-empty/okx_api/ws/public"
)

func main() {
	args := &ws.Args{
		Channel: "candle1m",
		InstId:  "BTC-USDT",
	}
	handler := func(c public.EventKline) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeKline(args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
