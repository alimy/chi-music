package core

import (
	"github.com/alimy/chi-music/models/model"
	"github.com/go-chi/render"
)

// MainAction indicator mina service interface
type MainAction interface {
	GetMainPage() (render.Renderer, error)
	AddAlbum(album *model.Album) error
	UpdateAlbum(album *model.Album) error
	GetAlbumById(id int64) (*model.Album, error)
	DeleteAlbumById(id int64) error
}
