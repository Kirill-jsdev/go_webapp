package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const portNumber = ":8080"





func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting app on port: %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}