package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	templates := populateTemplates() //Loading templates
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:] // removing leading / character
		//Find template based on requested page name
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			//Rendering and writing template to http.ResponseWriter which implements io.Writer
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			//Writing 404 to http.ResponseWriter.WriteHeader
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates") //returns a pointer to a new templates struct
	const basePath = "templates"        // templates directory
	//parsing all *.html matches in templates/ or fails
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
