package http

import (
	"net/http"

	"github.com/learninto/sniper-api/cmd/http/hooks"

	"github.com/learninto/goutil/twirp"
)

var commonHooks = twirp.ChainHooks(
	hooks.TraceID,
	hooks.Log,
)

func initMux(mux *http.ServeMux, isInternal bool) {

}

func initInternalMux(mux *http.ServeMux) {
}
