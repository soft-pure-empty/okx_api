package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples/rest"
	"github.com/soft-pure-empty/okx_api/rest/api"
	"github.com/soft-pure-empty/okx_api/rest/api/account"
)

func main() {
	param := &account.GetTradeFeeParam{
		InstId:   "BTC-USDT",
		InstType: api.InstTypeSPOT,
	}
	req, resp := account.NewGetTradeFee(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetTradeFeeResponse))
}
