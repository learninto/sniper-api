package server

import (
	twirphook "kingstar-go/sniper/twirp_hook"
	"net/http"

	"kingstar-go/sniper/twirp"
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

}

func initInternalMux(mux *http.ServeMux) {
}
