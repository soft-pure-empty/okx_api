package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples/rest"
	"github.com/soft-pure-empty/okx_api/rest/api/subaccount"
)

func main() {
	param := &subaccount.GetEntrustSubaccountListParam{}
	req, resp := subaccount.NewGetEntrustSubaccountList(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*subaccount.GetEntrustSubaccountListResponse))
}
