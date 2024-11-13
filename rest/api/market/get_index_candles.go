package market

import "github.com/soft-pure-empty/okx_api/rest/api"

func NewGetIndexCandles(param *GetCandlesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/index-candles",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCandlesResponse{}
}
