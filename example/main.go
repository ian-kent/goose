package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/ian-kent/goose"
)

var stream *goose.EventStream

func main() {
	stream = goose.NewEventStream()

	p := pat.New()
	p.Path("/events").Methods("GET").HandlerFunc(handler)
	p.Path("/notify").Methods("GET").HandlerFunc(someEvent)
	http.Handle("/", p)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func someEvent(w http.ResponseWriter, req *http.Request) {
	stream.Notify("data", []byte("Hi!"))
	w.WriteHeader(200)
}

func handler(w http.ResponseWriter, req *http.Request) {
	stream.AddReceiver(w)
}
