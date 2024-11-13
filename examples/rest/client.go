package rest

import (
	"github.com/soft-pure-empty/okx_api/examples"
	rc "github.com/soft-pure-empty/okx_api/rest"
)

// 敏感信息申请的模拟盘KEY，不确定何时会删除
var TestClient = rc.New("", examples.Auth, nil)
