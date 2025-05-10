package render

import (
	"bytes"
	"fmt"
	"log"
	"myapp/pkg/config"
	"net/http"
	"path/filepath"
	"text/template"
)
var app *config.AppConfig
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	if app.UseCache {

		//Get template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _  = CreateTemplateCache()
	}


	//get template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	//not mandatory stuff
	buff := new(bytes.Buffer)

	err := t.Execute(buff, nil)
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

func CreateTemplateCache() (map[string]*template.Template, error) {
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

