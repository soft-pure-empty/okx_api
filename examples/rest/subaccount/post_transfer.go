package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples/rest"
	"github.com/soft-pure-empty/okx_api/rest/api/subaccount"
)

func main() {
	param := &subaccount.PostTransferParam{
		Ccy:            "USDT",
		Amt:            "1",
		From:           "18",
		To:             "6",
		FromSubAccount: "test-01",
		ToSubAccount:   "test-02",
	}
	req, resp := subaccount.NewPostTransfer(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*subaccount.PostTransferResponse))
}
