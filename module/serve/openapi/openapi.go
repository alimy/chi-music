package openapi

import (
	"github.com/alimy/chi-music/models"
	"net/http"
)

// MirEntries get all entries that used to register to Mir
func MirEntries() []interface{} {
	ctx := models.NewContext()

	entries := []interface{}{
		&profile{Context: ctx},
		&media{Context: ctx},
	}
	if portal := mirPortal(); portal != nil {
		entries = append(entries, portal)
	}
	return entries
}

func httpResponse(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}
