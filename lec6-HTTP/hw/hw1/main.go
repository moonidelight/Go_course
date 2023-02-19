package main

import (
	"html/template"
	"net/http"
)

type ViewData struct {
	Title     string
	Paragraph string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title:     "About me",
			Paragraph: "Hello, my name is Zhanerke. I go to KBTU",
		}
		tmpl, _ := template.ParseFiles("about.html")
		tmpl.Execute(w, data)
	})
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		return
	}
}
