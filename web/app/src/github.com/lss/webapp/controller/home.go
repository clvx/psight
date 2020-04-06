package controller

import (
	"fmt"
	//"time"
	"html/template"
	"log"
	"net/http"

	"app/src/github.com/lss/webapp/viewmodel"
	"app/src/github.com/lss/webapp/model"
)

//home holds home.html template
type home struct {
	homeTemplate         *template.Template
	standLocatorTemplate *template.Template
	loginTemplate        *template.Template
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

	//triggering the timeout context to return a http.StatusRequestTimeout
	//time.Sleep(4*time.Second)

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
		if user, err := model.Login(email, password); err == nil {
			log.Printf("User has logged in: %v\n", user)
			//redirect to home
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			log.Printf("Failed to log user in with email: %v, error was: %v\n", email, err)
			vm.Email = email
			vm.Password = password
		}
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}
