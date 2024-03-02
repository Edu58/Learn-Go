package controllers

import (
	"log"
	"net/http"

	"github.com/Edu58/Learn-Go/WebChat/data"
	"github.com/Edu58/Learn-Go/WebChat/utils"
)

func CreateThread(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.GenerateHTML(w, nil, "base", "private.navbar", "create.thread")
	case "POST":
		err := r.ParseForm()

		if err != nil {
			log.Println(err)
			utils.SendError(w, r, "Something went wrong")
			return
		}

		user, err := utils.GetCurrentUser(w, r)

		if err != nil {
			utils.SendError(w, r, "Something Went Wrong")
		}

		thread := data.Thread{
			UserId: user.Uuid,
			Title:  r.PostFormValue("topic"),
		}

		err = thread.CreateThread()

		if err != nil {
			log.Println(err)
			utils.SendError(w, r, "Could not create thread")
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
