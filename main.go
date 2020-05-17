package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Dynamic template values
type HomeMetaData struct {
	Title string
	TagLine string
}

// Render homepage
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := HomeMetaData {
		Title: "HUSTLERS",
		TagLine: "NEW YORK'S FINEST GENTLEMAN CLUB.",
	}
	tmpl.Execute(w, data)
}

// Page routes
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	return r
}

// Init web server
func main() {
	router := Router()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9100",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}