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
		Name:     "",
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
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.PostFormValue("password")))

	if err != nil {
		log.Println(err)
	}

	if user.Email == r.PostFormValue("email") {
		session, err := user.CreateSession()
		log.Println(session)
		if err != nil {
			log.Println(err)
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		})
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
}
