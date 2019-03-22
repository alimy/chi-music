package openapi

import (
	"net/http"
)

// MirEntries get all entries that used to register to Mir
func MirEntries() []interface{} {
	entries := []interface{}{
		&profile{},
		&media{},
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
