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

	mux.HandleFunc("/signup", controllers.SignUp)
	mux.HandleFunc("/signin", controllers.SignIn)
	mux.HandleFunc("/signout", controllers.SignOut)

	mux.Handle("/", authMiddleware(http.HandlerFunc(controllers.Index)))
	mux.Handle("/thread", authMiddleware(http.HandlerFunc(controllers.Thread)))
	mux.Handle("/thread/new", authMiddleware(http.HandlerFunc(controllers.CreateThread)))
	mux.Handle("/err", authMiddleware(http.HandlerFunc(controllers.Err)))

	log.Print("Starting server on port 8000")
	log.Fatalln(http.ListenAndServe(":8000", mux))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		cookie, err := r.Cookie("_webchat")

		if err != nil || cookie == nil {
			log.Println("Error Getting Cookie")
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}

		session := data.Session{
			Uuid: cookie.Value,
		}

		valid, _ := session.Check()

		if valid {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}
	})
}
