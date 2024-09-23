package main

import "net/http"

func (a *application) heathCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
