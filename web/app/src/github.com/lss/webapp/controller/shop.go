package controller

import (
	"html/template"
	"net/http"
	"app/src/github.com/lss/webapp/viewmodel"
)

//holds shop.html template
type shop struct {
	shopTemplate *template.Template
}

//registerRoutes register the web routes
func (h shop) registerRoutes() {
	http.HandleFunc("/shop", h.handleShop)
}

//handleHome implements HandleFunc handler
func (h shop) handleShop(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewShop()
	//Rendering and writing template to io.Writer
	h.shopTemplate.Execute(w, vm)
}
