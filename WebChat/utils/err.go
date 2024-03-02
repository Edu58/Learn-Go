package utils

import (
	"net/http"
	"strings"
)

func SendError(w http.ResponseWriter, r *http.Request, msg string) {
	err := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(err, ""), http.StatusFound)
}
