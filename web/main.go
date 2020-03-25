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

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
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
	}(wg)
	wg.Wait()
}
