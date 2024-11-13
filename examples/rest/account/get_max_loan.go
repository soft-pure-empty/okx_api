package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples/rest"
	"github.com/soft-pure-empty/okx_api/rest/api"
	"github.com/soft-pure-empty/okx_api/rest/api/account"
)

func main() {
	param := &account.GetMaxLoanParam{
		InstId:  "BTC-USDT",
		MgnMode: api.MgnModeCross,
		MgnCcy:  "USDT",
	}
	req, resp := account.NewGetMaxLoan(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetMaxLoanResponse))
}
