package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"app/src/github.com/lss/webapp/controller"
)

func main() {
	templates := populateTemplates() //Loading templates
	controller.Startup(templates)
	
	http.ListenAndServe(":8000", nil)
}

//returns a map of strings to templates
func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	//Loading template
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	//Loading subtemplates
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	//Open content directory
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	//Readdir reads the contents of the directory associated with file and
	// returns a slice of up to n FileInfo values, as would be returned
	// by Lstat, in directory order.
	// If n <= 0, Readdir returns all the FileInfo from the directory in
	// a single slice
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	//Looping in all file info objects in contents/
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		//Reads template until error or EOF and returns data.
		// f implements io.Reader
		// ReadAll reads from r until an error or EOF and returns the data it read.
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()

		//Gets layout with all its children and clones it in a new template struct
		//Clone returns a duplicate of the template, including all associated
		// templates. The actual representation is not copied, but the name space of
		// associated templates is, so further calls to Parse in the copy will add
		// templates to the copy but not to the original.
		tmpl := template.Must(layout.Clone())
		//Creates template
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result

}
