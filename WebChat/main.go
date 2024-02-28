package main

import (
	"log"
	"net/http"

	"github.com/Edu58/Learn-Go/WebChat/controllers"
	"github.com/Edu58/Learn-Go/WebChat/data"
)

func init() {
	data.DBConn()
}
func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.Handle("/", authMiddleware(http.HandlerFunc(controllers.Index)))
	mux.HandleFunc("/signup", controllers.SignUp)
	mux.HandleFunc("/signin", controllers.SignIn)
	mux.HandleFunc("/signout", controllers.SignOut)

	log.Print("Starting server on port 8000")
	log.Fatalln(http.ListenAndServe(":8000", mux))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		cookie, err := r.Cookie("_cookie")
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/signin", http.StatusFound)
		}

		session := data.Session{
			Uuid: cookie.Value,
		}

		valid, _ := session.Check()

		if valid {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/signin", http.StatusFound)
		}
	})
}
