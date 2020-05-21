package server

import (
	"kingstar-go/sniper/trace"
	twirphook "kingstar-go/sniper/twirp_hook"
	"net/http"

	"github.com/bilibili/twirp"
)

var hooks = twirp.ChainHooks(
	twirphook.NewHeaders(),
	trace.NewRequestID(),
	twirphook.NewLog(),
)

var loginHooks = twirp.ChainHooks(
	twirphook.NewHeaders(),
	trace.NewRequestID(),
	twirphook.NewLog(),
	twirphook.NewCheckLogin(),
)

func initMux(mux *http.ServeMux, isInternal bool) {

}

func initInternalMux(mux *http.ServeMux) {
}
