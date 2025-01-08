package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	/*
			To access multiple form uploads

		fhs := req.MultipartForm.File["myfiles"]
		for _, fh := range fhs {
		    f, err := fh.Open()
		    // f is one of the files
		}
	*/
	err := r.ParseMultipartForm(32 << 20)

	if err != nil {
		log.Fatalln(err)
	}

	file, handler, err := r.FormFile("file")

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	dst, err := os.Create(handler.Filename)

	if err != nil {
		log.Fatalln(err)
	}

	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		log.Fatalln(err)
	}

	w.WriteHeader(http.StatusOK)
}
