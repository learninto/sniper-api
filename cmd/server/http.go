package server

import (
	"kingstar-go/sniper/trace"
	twirphook "kingstar-go/sniper/twirp-hook"
	"net/http"

	"sniper-api/cmd/server/hook"

	"github.com/bilibili/twirp"

	shop_v1 "sniper-api/rpc/shop/v1"
	"sniper-api/server/shopserver1"
)

var hooks = twirp.ChainHooks(
	trace.NewRequestID(),
	twirphook.NewLog(),
)

var loginHooks = twirp.ChainHooks(
	trace.NewRequestID(),
	twirphook.NewLog(),
	hook.NewCheckLogin(),
)

func initMux(mux *http.ServeMux, isInternal bool) {
	{
		server := &shopserver1.Server{}
		handler := shop_v1.NewShopServer(server, hooks)
		mux.Handle(shop_v1.ShopPathPrefix, handler)
	}
}

func initInternalMux(mux *http.ServeMux) {
}
