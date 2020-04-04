package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"app/src/github.com/lss/webapp/viewmodel"
)

//home holds home.html template
type home struct {
	homeTemplate         *template.Template
	standLocatorTemplate *template.Template
	loginTemplate		 *template.Template
}

//registerRoutes register the web routes
func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/stand-locator", h.handleStandLocator)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	w.Header().Add("Content-Type", "text/html")
	//Rendering and writing template to io.Writer
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewStandLocator()
	h.standLocatorTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()
	//checking if the request is POST
	if r.Method == http.MethodPost {
		err := r.ParseForm() //parsing form
		if err != nil {
			log.Println(fmt.Errorf("Error logging in: %v", err))
		}
		email := r.Form.Get("email")       //getting email from form
		password := r.Form.Get("password") //getting email from form
		if email == "test@gmail.com" && password == "password" {
			//redirect to home
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			vm.Email = email
			vm.Password = password
		}
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}
