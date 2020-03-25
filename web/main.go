package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sync"
)

/*
Implementing the Handle interface
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
type myHandler struct {
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

const tax = 6.75 / 100

type Product struct {
	Name  string
	Price float32
}

func (p Product) PriceWithTax() float32 {
	return p.Price * (1 + tax)
}

const templateString = `
{{- "Item information" }}
Name: {{ .Name }}
Price: {{ printf "$%.2f" .Price }}
Price with Tax: {{ .PriceWithTax | printf "$%.2f" }} //This is a method, same formatting as the previous line
`

const customTemplateString = `
{{- "Item Information" }}
Name: {{ .Name }}
Price: {{ printf "$%.2f" .Price }}
Price with Tax: {{ calctax .Price | printf "$%.2f" }}
`

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go func(wg *sync.WaitGroup) {
		//Handle registers the handler for the given pattern
		// in the DefaultServerMux
		//func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }
		// pattern: go wil use the handler which has the most specific match
		http.Handle("/foo", &myHandler{greeting: "Hello"})
		http.ListenAndServe(":8000", nil)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		// func HandleFunc(pattern string, handler HandlerFunc)
		//Keep in mind `handler HandlerFunc` appears in `server.go` as func(ResponseWriter, *Request)
		// type HandlerFunc func(ResponseWriter, *Request)
		// func (f HandleFunc) ServeHTTP(w ResponseWriter, r *Request)
		http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World"))
		})
		http.ListenAndServe(":8001", nil)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		http.HandleFunc("/baz", func(w http.ResponseWriter, r *http.Request) {
			templateString := `Lemonade Stand Supply`
			t, err := template.New("title").Parse(templateString)
			if err != nil {
				fmt.Println(err)
			}
			//Execute applies a parsed template to the specified data object,
			// writing the output to wr.
			//os.Stdout implements io.Writer interface
			err = t.Execute(os.Stdout, nil)
			if err != nil {
				fmt.Println(err)
			}
			//w implementes io.Writer interface.
			err = t.Execute(w, nil)
			if err != nil {
				fmt.Println(err)
			}
		})
		http.ListenAndServe(":8002", nil)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		http.HandleFunc("/foo/price", func(w http.ResponseWriter, r *http.Request) {

			p := Product{
				Name:  "Lemonade",
				Price: 2.16,
			}
			t := template.Must(template.New("").Parse(templateString))
			t.Execute(w, p)
		})
		http.ListenAndServe(":8003", nil)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		http.HandleFunc("/foo/custom-price", func(w http.ResponseWriter, r *http.Request) {
			p := Product{
				Name:  "Lemonade",
				Price: 2.16,
			}
			fm := template.FuncMap{} //Creating an object of type FuncMap
			//Creating a custom function instead of using methods as the previous example
			// FuncMap expects a map of interfaces which accepts any type as a value.
			fm["calctax"] = func(price float32) float32 {
				return price * (1 + tax)
			}
			t := template.Must(template.New("").Funcs(fm).Parse(customTemplateString))
			t.Execute(w, p)
		})
		http.ListenAndServe(":8004", nil)
		wg.Done()
	}(wg)

	wg.Wait()
}
