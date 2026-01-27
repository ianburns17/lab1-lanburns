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
func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact me at:\n email: ianburns344@gmail.com \n phone number: 6259908 "))
}
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
