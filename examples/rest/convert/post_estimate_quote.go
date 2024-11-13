package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples/rest"
	"github.com/soft-pure-empty/okx_api/rest/api"
	"github.com/soft-pure-empty/okx_api/rest/api/convert"
)

func main() {
	param := &convert.PostEstimateQuoteParam{
		BaseCcy:  "BTC",
		QuoteCcy: "USDT",
		Side:     api.SideSell,
		RfqSz:    "10",
		RfqSzCcy: "USDT",
	}
	req, resp := convert.NewPostEstimateQuote(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.PostEstimateQuoteResponse))
}
