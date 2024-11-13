package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples/rest"
	"github.com/soft-pure-empty/okx_api/rest/api/market"
)

func main() {
	param := &market.GetCandlesParam{
		InstId: "ETH-USDT",
	}
	req, resp := market.NewGetCandles(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetCandlesResponse))
}
