package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello sir, My name is Ian burns. I'll mostly choose the car rental system due to its complexity and usability. Also it would be useful for a family member who owns a car rental.\n"))
}
func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My name is Ian Burns, and I am an 18-year-old with a passion for cybersecurity. In my free time, I enjoy going on adventures and trying new things.n"))
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact me at:\n email: ianburns344@gmail.com \n phone number: 6259908 "))
}
func qoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("“The best way to get started is to quit talking and begin doing.” – Walt Disney\n"))
}
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/qoute", qoute)
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
