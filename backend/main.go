package main

import (
	"encoding/json"
	"io"

	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		fmt.Println(string(d))

		fmt.Println("hello world")

		type result struct {
			Hello string `json:"hello"`
		}

		b, err := json.Marshal(result{Hello: "world"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(string(b))

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

		fmt.Println("hello world")
	})

	http.ListenAndServe(":8080", nil)
}
