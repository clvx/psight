package controller

import (
	"html/template"
	"net/http"

	"app/src/github.com/lss/webapp/viewmodel"
)

//home holds home.html template
type home struct {
	homeTemplate *template.Template
}

//registerRoutes register the web routes
func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
}

//handleHome implements HandleFunc handler
func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	//Rendering and writing template to io.Writer
	h.homeTemplate.Execute(w, vm)
}
