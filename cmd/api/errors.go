package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err)

	writeInternalError(w)
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("bad request error", "method", r.Method, "path", r.URL.Path, "error", err)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundErorr(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("not found error", "method", r.Method, "path", r.URL.Path, "error", err)

	writeJSONError(w, http.StatusNotFound, err.Error())
}

func (app *application) conflictError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("conflict error", "method", r.Method, "path", r.URL.Path, "error", err)

	writeJSONError(w, http.StatusConflict, err.Error())
}
