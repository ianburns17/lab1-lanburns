package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from newsletter\n"))
}
func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About me- My name is Ian Burns and i Love cyber security\n"))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
