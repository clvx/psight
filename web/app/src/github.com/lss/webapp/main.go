package main

import (
	"net/http"
)

func main() {
	/*
	//This piece of code gives the same functionality as the one liner below
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// ServeFile replies to the request with the contents of the named
		// file or directory.
		http.ServeFile(w, r, "public"+r.URL.Path)
		
	})
	*/

	//FileServer returns a handler that serves HTTP requests
	// with the contents of the file system rooted at root.
	//A Dir implements FileSystem using the native file system restricted to a
	// specific directory tree.
	http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}
