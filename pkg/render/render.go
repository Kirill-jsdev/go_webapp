package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	//Create template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//get template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	//not mandatory stuff
	buff := new(bytes.Buffer)

	err = t.Execute(buff, nil)
	if err != nil {
		fmt.Println(err)
	}
	//end

	//render template
	_, err = buff.WriteTo(w)

	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	//get all pages from a certain folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page) //returns the last element of the path

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, nil
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, nil
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, nil
			}
		}

		myCache[name] = ts
	}

	return myCache, nil

}

