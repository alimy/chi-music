package openapi

import (
	"net/http"
)

func httpResponse(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}
