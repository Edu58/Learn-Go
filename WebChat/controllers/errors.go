package controllers

import (
	"net/http"

	"github.com/Edu58/Learn-Go/WebChat/utils"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/signin", http.StatusFound)
}

func Err(w http.ResponseWriter, r *http.Request) {
	err_msg := r.URL.Query().Get("msg")
	utils.GenerateHTML(w, err_msg, "base", "private.navbar", "error")
}
