package server

import (
	"net/http"

	twirphook "github.com/learninto/sniper-api/utils/twirp_hook"

	demo_v1 "github.com/learninto/sniper-api/rpc/demo/v1"
	"github.com/learninto/sniper-api/utils/twirp"
)

var hooks = twirp.ChainHooks(
	twirphook.NewHeaders(),
	twirphook.NewRequestID(),
	twirphook.NewLog(),
)

var loginHooks = twirp.ChainHooks(
	twirphook.NewHeaders(),
	twirphook.NewRequestID(),
	twirphook.NewLog(),
	twirphook.NewCheckLogin(),
)

func initMux(mux *http.ServeMux, isInternal bool) {

	{
		server := &demo_v1.DemoServer{}
		handler := demo_v1.NewDemoServer(server, hooks)
		mux.Handle(demo_v1.DemoPathPrefix, handler)
	}
}

func initInternalMux(mux *http.ServeMux) {
}
