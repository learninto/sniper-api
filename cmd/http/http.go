package http

import (
	"net/http"

	twirphook "github.com/learninto/goutil/twirp_hook"

	"github.com/learninto/goutil/twirp"
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
