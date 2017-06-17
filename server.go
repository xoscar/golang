package main

import (
	"net/http"
	"encoding/json"
	// "io"
	// "fmt"
)

type Human struct {
	Age int `json:"age"`
	Name  string `json:"name"`
}

type Crowd struct {
	People []Human
	Leader Human 
}

func main() {
	// http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, r.URL.Path)
	// })

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("views"))))

	http.HandleFunc("/json", func (w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(
			Crowd {
				People: []Human{{23, "Oscar"}, {30, "Carlos"}},
				Leader: Human{20, "Juan"},
			},
		)
	})

	http.ListenAndServe(":8080", nil)
}