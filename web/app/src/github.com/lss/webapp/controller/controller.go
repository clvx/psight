package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home //home type
	shopController shop //shop type
)

func Startup(templates map[string]*template.Template) {
	//assign templates
	homeController.homeTemplate = templates["home.html"]
	shopController.shopTemplate = templates["shop.html"]
	//register routes
	homeController.registerRoutes()
	shopController.registerRoutes()
	//register statics
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
