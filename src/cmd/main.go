package main

import (
	"fizzbuzz/fizzbuzz"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/fizzbuzz", func(w http.ResponseWriter, r *http.Request) {
		number := r.URL.Query().Get("number")
		w.Write([]byte(fizzbuzz.FizzBuzz(number)))
	})
	http.ListenAndServe(port, nil)
}
