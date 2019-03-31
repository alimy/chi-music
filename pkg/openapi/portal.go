// +build portal

package openapi

import (
	"github.com/alimy/mir"
	"github.com/alimy/music-ui/dist"
	"net/http"
)

type portalAssets struct {
	index            mir.Get  `mir:/index.html#GetMainAssets`
	getMainAssets    mir.Get  `mir:"/"`
	getStaticAssets  mir.Get  `mir:"/static/*"`
	headStaticAssets mir.Head `mir:"/static/*"`

	staticHandler http.Handler
}

// GetMainAssets GET handler of "/"
func (p *portalAssets) GetMainAssets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(dist.MustAsset("index.html"))
}

// GetStaticAssets GET handler of "/static/*filepath"
func (p *portalAssets) GetStaticAssets(w http.ResponseWriter, r *http.Request) {
	p.staticHandler.ServeHTTP(w, r)
}

// HeadStaticAssets HEAD handler of "/static/*filepath"
func (p *portalAssets) HeadStaticAssets(w http.ResponseWriter, r *http.Request) {
	p.staticHandler.ServeHTTP(w, r)
}

// MirPortal return a portal mir entry
func MirPortal() interface{} {
	assetFile := dist.AssetFile()
	return &portalAssets{
		staticHandler: http.StripPrefix("/static", http.FileServer(assetFile)),
	}
}
