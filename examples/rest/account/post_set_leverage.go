package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples/rest"
	"github.com/soft-pure-empty/okx_api/rest/api"
	"github.com/soft-pure-empty/okx_api/rest/api/account"
)

func main() {
	param := &account.PostSetLeverageParam{
		InstId:  "BTC-USDT",
		Lever:   "5",
		MgnMode: api.MgnModeCross,
	}
	req, resp := account.NewPostSetLeverage(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.PostSetLeverageResponse))
}
