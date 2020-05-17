package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

type HomeMetaData struct {
	Title string
	Tagline string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := HomeMetaData {
		Title: "HUSTLERS",
		Tagline: "NEW YORK'S FINEST GENTLEMAN CLUB.",
	}
	tmpl.Execute(w, data)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	//r.PathPrefix("/").Handler(spa)
	return r
}

func main() {
	router := newRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9100",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}