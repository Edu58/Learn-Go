package controllers

import (
	"log"
	"net/http"

	"github.com/Edu58/Learn-Go/WebChat/data"
	"github.com/Edu58/Learn-Go/WebChat/utils"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.GenerateHTML(w, nil, "base", "signin", "public.navbar")
	case "POST":
		authSignIn(w, r)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.GenerateHTML(w, nil, "base", "signup", "public.navbar")
	case "POST":
		authSignUp(w, r)
	}
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Fatalln(err)
		return
	}

	session := data.Session{
		Uuid: cookie.Value,
	}

	session.DeleteSession()

	http.Redirect(w, r, "/signin", http.StatusFound)
}
