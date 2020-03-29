package controller

import (
	"html/template"
	"net/http"

	"app/src/github.com/lss/webapp/viewmodel"
)

//home holds home.html template
type home struct {
	homeTemplate         *template.Template
	standLocatorTemplate *template.Template
}

//registerRoutes register the web routes
func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/stand-locator", h.handleStandLocator)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	//Rendering and writing template to io.Writer
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewStandLocator()
	h.standLocatorTemplate.Execute(w, vm)
}
