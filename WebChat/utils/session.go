package utils

import (
	"log"
	"net/http"

	"github.com/Edu58/Learn-Go/WebChat/data"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request) (user data.User, err error) {
	cookie, error := r.Cookie("_webchat")

	if error != nil {
		SendError(w, r, "Something went wrong")
		return
	}

	session := data.Session{
		Uuid: cookie.Value,
	}

	current_session, err := session.GetSessionById()

	if err != nil {
		log.Fatalln("Error")
		return
	}

	user, err = data.GetUserById(current_session.UserId)

	if err != nil {
		log.Fatalln("Error")
		SendError(w, r, "Something went wrong")
	}
	return
}
