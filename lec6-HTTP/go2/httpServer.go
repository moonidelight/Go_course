package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/about", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "about page")
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "contact page")
	})
	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		return
	}
}
