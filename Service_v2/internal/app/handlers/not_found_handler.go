package handlers

import (
	"errors"
	"net/http"
)

func NotFound(w http.ResponseWriter, _ *http.Request) {
	WrapErrorWithStatus(w, errors.New("not found"), http.StatusNotFound)
}
