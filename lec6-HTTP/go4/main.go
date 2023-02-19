package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "user.html")
	})
	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("username")
		age := r.FormValue("userage")
		fmt.Fprintf(w, "Name: %s; Age: %s", name, age)
	})
	fmt.Println("Server is listening")
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		return
	}
}
