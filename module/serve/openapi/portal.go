// +build portal

package openapi

import (
	"github.com/alimy/mir"
	"github.com/alimy/music-ui/dist"
	"github.com/elazarl/go-bindata-assetfs"
	"net/http"
)

type portalAssets struct {
	mainHandler      http.Handler
	staticHandler    http.Handler
	getMainAssets    mir.Get  `mir:"/"`
	headMainAssets   mir.Head `mir:"/"`
	getStaticAssets  mir.Get  `mir:"/static/*"`
	headStaticAssets mir.Head `mir:"/static/*"`
}

// GetMainAssets GET handler of "/"
func (p *portalAssets) GetMainAssets(w http.ResponseWriter, r *http.Request) {
	p.mainHandler.ServeHTTP(w, r)
}

// HeadMainAssets HEAD handler of "/"
func (p *portalAssets) HeadMainAssets(w http.ResponseWriter, r *http.Request) {
	p.mainHandler.ServeHTTP(w, r)
}

// GetStaticAssets GET handler of "/static/*filepath"
func (p *portalAssets) GetStaticAssets(w http.ResponseWriter, r *http.Request) {
	p.staticHandler.ServeHTTP(w, r)
}

// HeadStaticAssets HEAD handler of "/static/*filepath"
func (p *portalAssets) HeadStaticAssets(w http.ResponseWriter, r *http.Request) {
	p.staticHandler.ServeHTTP(w, r)
}

// mirPortal return a portal mir entry
func mirPortal() interface{} {
	assetFS := &assetfs.AssetFS{
		Asset:     dist.Asset,
		AssetDir:  dist.AssetDir,
		AssetInfo: dist.AssetInfo}
	return &portalAssets{
		mainHandler:   http.StripPrefix("/", http.FileServer(assetFS)),
		staticHandler: http.StripPrefix("/static", http.FileServer(assetFS)),
	}
}