package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello, world!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes written: %d\n", n)
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "About!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes written: %d\n", n)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting app on port: %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}