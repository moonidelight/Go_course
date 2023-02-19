package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")
		fmt.Fprintf(w, "name: %s; age: %s", name, age)
	})

	fmt.Println("server is listening")
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		return
	}
}
