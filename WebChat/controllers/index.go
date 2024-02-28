package controllers

import (
	"net/http"

	"github.com/Edu58/Learn-Go/WebChat/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.GenerateHTML(w, nil, "base", "index", "private.navbar")
}
