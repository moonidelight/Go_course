package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Fprintf(w, "<h1>Time: %s</h1>", t)
	})
	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		return
	}
}
