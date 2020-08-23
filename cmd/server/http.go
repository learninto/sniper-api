package server

import (
	"net/http"

	twirphook "github.com/learninto/sniper-api/utils/twirp_hook"

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

}

func initInternalMux(mux *http.ServeMux) {
}
