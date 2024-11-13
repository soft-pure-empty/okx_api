package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples/rest"
	"github.com/soft-pure-empty/okx_api/rest/api/convert"
)

func main() {
	param := &convert.PostTradeParam{
		BaseCcy:  "BTC",
		QuoteCcy: "USDT",
		Side:     "buy",
		Sz:       "1",
		SzCcy:    "USDT",
		QuoteId:  "123456789",
	}
	req, resp := convert.NewPostTrade(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.PostTradeResponse))
}
