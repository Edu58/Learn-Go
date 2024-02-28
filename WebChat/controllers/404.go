package controllers

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/signin", http.StatusFound)
}
