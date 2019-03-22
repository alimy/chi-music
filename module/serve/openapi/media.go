package openapi

import (
	"github.com/alimy/chi-music/models"
	"github.com/alimy/chi-music/models/core"
	"github.com/alimy/mir"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/unisx/logus"
	"net/http"
)

type media struct {
	group            mir.Group  `mir:"v1"`
	getAlbums        mir.Get    `mir:"/albums"`
	createAlbums     mir.Put    `mir:"/albums"`
	updateAlbums     mir.Post   `mir:"/albums"`
	getAlbumsById    mir.Get    `mir:"/albums/{albumId}"`
	deleteAlbumsById mir.Delete `mir:"/albums/{albumId}"`
}

// GetAlbums GET handler of "/albums/"
func (m *media) GetAlbums(w http.ResponseWriter, r *http.Request) {
	if albums, ok := core.Model(models.IdAlbums).(*models.Albums); ok {
		logus.Debug("getAlbums")
		core.Retrieve(models.IdAlbums, albums)
		render.Status(r, http.StatusOK)
		render.Render(w, r, albums)
	} else {
		httpResponse(w, http.StatusNotFound, "albums not found")
	}
}

// CreateAlbums PUT handler of "/albums/"
func (m *media) CreateAlbums(w http.ResponseWriter, r *http.Request) {
	// TODO
	logus.Debug("create albums")
	httpResponse(w, http.StatusCreated, "albums item created")
}

// UpdateAlbums POST handler of "/albums/"
func (m *media) UpdateAlbums(w http.ResponseWriter, r *http.Request) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("updateAlbums")
		album.Id = chi.URLParam(r, "albumId")
		core.Update(models.IdAlbum, album)
		httpResponse(w, http.StatusCreated, "albums item updated")
	} else {
		httpResponse(w, http.StatusNotFound, "update albums failure")
	}
}

// GetAlbumsById GET handler of "/albums/:albumId/"
func (m *media) GetAlbumsById(w http.ResponseWriter, r *http.Request) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("getAlbumsById")
		album.Id = chi.URLParam(r, "albumId")
		core.Retrieve(models.IdAlbum, album)
		render.Render(w, r, album)
	} else {
		httpResponse(w, http.StatusNotFound, "update albums failure")
	}
}

// DeleteAlbumsById DELETE handler of "/albums/:albumId/"
func (m *media) DeleteAlbumsById(w http.ResponseWriter, r *http.Request) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("deleteAlbumsById")
		album.Id = chi.URLParam(r, "albumId")
		core.Delete(models.IdAlbum, album)
		httpResponse(w, http.StatusOK, "albums item deleted")
	} else {
		httpResponse(w, http.StatusNotFound, "delete album failure")
	}
}
