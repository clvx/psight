package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//look into the request object, check the file, and translate it to the fs
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()

		//Check request content type
		var contentType string
		switch {
		case strings.HasSuffix(r.URL.Path, "css"):
			contentType = "text/css"
		case strings.HasSuffix(r.URL.Path, "html"):
			contentType = "text/html"
		case strings.HasSuffix(r.URL.Path, "png"):
			contentType = "text/png"
		default:
			contentType = "text/plain"
		}
		//Add a Content-Type header in ResponseWriter object
		w.Header().Add("Content-Type", contentType)
		//Copy file content directly to http.ResponseWriter which implements Writer interface
		io.Copy(w, f)
	})
	http.ListenAndServe(":8000", nil)
}
