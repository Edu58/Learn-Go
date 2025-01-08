package utils

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(createdAt time.Time) string {
	return createdAt.Format("Jan 2, 2006 at 3:04pm")
}

func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	for _, file := range filenames {
		files = append(files, "templates/"+file+".html")
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
