package main

import (
	"log"

	"github.com/soft-pure-empty/okx_api/examples/rest"
	"github.com/soft-pure-empty/okx_api/rest/api/asset"
)

func main() {
	param := &asset.GetBalancesParam{}
	req, resp := asset.NewGetBalances(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*asset.GetBalancesResponse))
}
