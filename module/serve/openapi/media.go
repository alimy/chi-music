package openapi

import (
	"encoding/json"
	"github.com/alimy/chi-music/models/core"
	"github.com/alimy/chi-music/models/model"
	"github.com/alimy/mir"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/unisx/logus"
	"net/http"
	"strconv"
)

type media struct {
	group            mir.Group  `mir:"v1"`
	getAlbums        mir.Get    `mir:"/albums"`
	createAlbums     mir.Put    `mir:"/albums"`
	updateAlbums     mir.Post   `mir:"/albums"`
	getAlbumsById    mir.Get    `mir:"/albums/{albumId}"`
	deleteAlbumsById mir.Delete `mir:"/albums/{albumId}"`

	*core.Context
}

// GetAlbums GET handler of "/albums/"
func (m *media) GetAlbums(w http.ResponseWriter, r *http.Request) {
	m.Retrieve(w, r, core.RdsMainPage, m.Repo.GetMainPage)
}

// CreateAlbums PUT handler of "/albums/"
func (m *media) CreateAlbums(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	album := &model.Album{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(album)

	if err := m.Repo.AddAlbum(album); err != nil {
		m.ErrInternalServer(w, err.Error())
		logus.Debug("create albums failure", logus.ErrorField(err))
		return
	}
	m.Expire(core.RdsMainPage)
	logus.Debug("create albums success")
	m.HttpResponse(w, http.StatusCreated, "albums item created")
}

// UpdateAlbums POST handler of "/albums/"
func (m *media) UpdateAlbums(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	album := &model.Album{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(album)

	if album.AlbumId == 0 {
		if err := m.Repo.AddAlbum(album); err != nil {
			m.ErrInternalServer(w, err.Error())
			logus.Debug("create albums failure", logus.ErrorField(err))
			return
		}
		m.Expire(core.RdsMainPage)
		logus.Debug("create albums success")
		m.HttpResponse(w, http.StatusCreated, "albums item created")
		return
	}

	if err := m.Repo.UpdateAlbum(album); err != nil {
		m.ErrInternalServer(w, err.Error())
		logus.Debug("update albums failure", logus.ErrorField(err))
		return
	}
	m.Expire(core.RdsMainPage)
	logus.Debug("update albums success")
	m.HttpResponse(w, http.StatusCreated, "albums item created")
}

// GetAlbumsById GET handler of "/albums/{albumId}"
func (m *media) GetAlbumsById(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	id, err := strconv.ParseInt(albumId, 10, 0)
	if err != nil {
		m.ErrInternalServer(w, err.Error())
		logus.Debug("parse id failure", logus.ErrorField(err))
		return
	}
	album, err := m.Repo.GetAlbumById(id)
	if err != nil {
		m.ErrInternalServer(w, err.Error())
		logus.Debug("get album by id failure", logus.Int64("id", id), logus.ErrorField(err))
		return
	}
	err = render.Render(w, r, album)
	if err != nil {
		m.ErrInternalServer(w, err.Error())
		logus.Debug("render album failure", logus.Int64("id", id), logus.ErrorField(err))
		return
	}
	logus.Debug("get albums by id success", logus.Int64("id", id))
}

// DeleteAlbumsById DELETE handler of "/albums/{albumId}"
func (m *media) DeleteAlbumsById(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	id, err := strconv.ParseInt(albumId, 10, 0)
	if err != nil {
		m.ErrInternalServer(w, err.Error())
		logus.Debug("parse id failure", logus.ErrorField(err))
		return
	}
	err = m.Repo.DeleteAlbumById(id)
	if err != nil {
		m.ErrInternalServer(w, err.Error())
		logus.Debug("delete album failure", logus.Int64("id", id), logus.ErrorField(err))
		return
	}
	m.Expire(core.RdsMainPage)
	logus.Debug("delete album by id success", logus.Int64("id", id))
	m.HttpResponse(w, http.StatusOK, "album deleted")
}
