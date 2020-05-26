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
	Title   string
	TagLine string
	SiteUrl string
	ShareImage string
	Icon string
}

// Render homepage
func indexHandler (w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := HomeMetaData{
		Title:   "HUSTLERS",
		TagLine: "NEW YORK'S FINEST GENTLEMAN CLUB.",
		SiteUrl: "https://hustlers.club/",
		ShareImage: "https://hackers-content.nyc3.digitaloceanspaces.com/sites/hustlers/img/hustlersshare@2x.jpg",
		Icon: "https://hackers-content.nyc3.digitaloceanspaces.com/sites/hustlers/img/icon.png",
	}
	tmpl.Execute(w, data)
}

// Route declaration
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	return r
}

// Initiate web server
func main() {
	router := Router()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}