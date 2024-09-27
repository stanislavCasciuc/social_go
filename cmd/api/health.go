package main

import (
	"net/http"
)

func (a *application) heathCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"status": "OK"}
	if err := WriteJSON(w, http.StatusOK, data); err != nil {
		writeInternalError(w)
	}
}
