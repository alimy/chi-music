package openapi

import (
	"github.com/alimy/chi-music/models/core"
	"github.com/alimy/mir"
	"github.com/unisx/logus"
	"net/http"
)

type profile struct {
	group      mir.Group `mir:"v1"`
	getAppInfo mir.Get   `mir:"/appinfo"`

	*core.Context
}

// GetAppInfo GET handler of "/appinfo/"
func (p *profile) GetAppInfo(w http.ResponseWriter, r *http.Request) {
	// TODO
	logus.Debug("get application information")
	httpResponse(w, http.StatusOK, "get application information")
}
