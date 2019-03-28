package core

import (
	"github.com/alimy/chi-music/models/cache"
	"github.com/go-chi/render"
	"net/http"
)

// Context core context for gin handler
type Context struct {
	cache.Cache

	Repo Repository
}

// Retrieve write response content to c
func (ctx *Context) Retrieve(w http.ResponseWriter, r *http.Request, cacheKey string, action func() (render.Renderer, error)) {
	if ok := ctx.Cache.EntryTo(cacheKey, w); !ok {
		data, err := action()
		if err == nil {
			err = render.Render(w, r, data)
			ctx.CacheFrom(cacheKey, data)
		}
		if err != nil {
			ctx.ErrInternalServer(w, err.Error())
		}
	}
}

func (ctx *Context) ErrInternalServer(w http.ResponseWriter, msg ...string) {
	w.WriteHeader(http.StatusInternalServerError)
	if len(msg) > 1 {
		w.Write([]byte(msg[0]))
	} else {
		w.Write([]byte("internal server error"))
	}
}

func (ctx *Context) ErrNotFound(w http.ResponseWriter, msg ...string) {
	w.WriteHeader(http.StatusNotFound)
	if len(msg) > 1 {
		w.Write([]byte(msg[0]))
	} else {
		w.Write([]byte("not found"))
	}
}

func (ctx *Context) HttpResponse(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}
