package controllers

import (
	"log"
	"net/http"

	"github.com/Edu58/Learn-Go/WebChat/data"
	"github.com/Edu58/Learn-Go/WebChat/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.GetThreads()

	if err != nil {
		log.Fatalln(err)
		utils.SendError(w, r, "Could not get threads")
	}

	utils.GenerateHTML(w, threads, "base", "index", "private.navbar")
}
