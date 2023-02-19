package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := "New Template"
		tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1>")
		tmpl.Execute(w, data)
	})
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		return
	}
}
