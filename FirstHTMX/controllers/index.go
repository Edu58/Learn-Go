package controllers

import (
	"html/template"
	"net/http"

	"github.com/Edu58/Learn-Go/FirstHTMX/utils"
)

type Film struct {
	Name     string
	Director string
}

var films = map[string][]Film{
	"Films": {
		{"The Matrix", "The Wachowskis"},
		{"The Matrix Reloaded", "The Wachowskis"},
		{"The Matrix Revolutions", "The Wachowskis"},
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	utils.GenerateHTML(w, films, "base", "index")
}

func Create(w http.ResponseWriter, r *http.Request) {
	film := Film{
		Name:     r.PostFormValue("name"),
		Director: r.PostFormValue("director"),
	}
	tmpl, _ := template.New("create").ParseFiles("templates/index.html")
	tmpl.ExecuteTemplate(w, "film-list", film)
}
