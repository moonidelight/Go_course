package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "info.html")
	})
	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		first := r.FormValue("16-digits")
		second := r.FormValue("cvv")
		if len(first) == 16 && len(second) == 3 {
			//io.WriteString(w, "Card is valid")
			fmt.Fprint(w, "Card is valid")
			return
		}
		//io.WriteString(w, "Card is not valid")
		fmt.Fprint(w, "Not valid card")
	})
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		return
	}
}
