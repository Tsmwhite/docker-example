package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hey boy"))
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
