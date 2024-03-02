package controllers

import (
	"log"
	"net/http"

	"github.com/Edu58/Learn-Go/WebChat/data"
	"golang.org/x/crypto/bcrypt"
)

func authSignUp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Println(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), 14)

	if err != nil {
		log.Fatalln(err)
		return
	}

	user := &data.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: string(hashedPassword),
	}

	if err := user.CreateUser(); err != nil {
		log.Fatalln(err)
		http.Redirect(w, r, "/sigup", http.StatusFound)
	}

	http.Redirect(w, r, "/signin", http.StatusFound)
}

func authSignIn(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Println(err)
	}

	user, err := data.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.PostFormValue("password")))

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	if user.Email == r.PostFormValue("email") {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := &http.Cookie{
			Name:     "_webchat",
			Value:    session.Uuid,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
}
