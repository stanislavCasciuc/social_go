package main

import (
	"net/http"
)

// healthcheckHandler godoc
//
//	@Summary		Healthcheck
//	@Description	Healthcheck endpoint
//	@Tags			ops
//	@Produce		json
//	@Success		200	{object}	string	"ok"
//	@Router			/health [get]
func (a *application) heathCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"status": "OK"}
	if err := WriteJSON(w, http.StatusOK, data); err != nil {
		writeInternalError(w)
	}
}
