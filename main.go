package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kib357/less-go"
	"html/template"
	"log"
	"net/http"
	"time"
)

// HomeMetaData Dynamic template values
type HomeMetaData struct {
	Title      string
	TagLine    string
	FooterText string
	SiteUrl    string
	ShareImage string
	Icon       string
}

// Render homepage
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := HomeMetaData{
		Title:      "HUSTLERS",
		TagLine:    "NEW YORK'S FINEST GENTLEMAN'S CLUB",
		FooterText: "A Broiest Production™",
		SiteUrl:    "https://hustlers.club/",
		ShareImage: "https://storage.cloud.google.com/hustlers/img/hustlersshare@2x.jpg",
		Icon:       "https://storage.cloud.google.com/hustlers/img/icon.png",
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		return
	}
}


// CompileStylesheets Compile and minify .LESS files
func CompileStylesheets() {
	staticFolder := "./static/%s"
	err := less.RenderFile(
		fmt.Sprintf(staticFolder, "style.less"),
		fmt.Sprintf(staticFolder, "style.css"),
		map[string]interface{}{"compress": true})
	if err != nil {
		log.Fatal(err)
	}
}

// Router Route declaration
func Router() *mux.Router {
	staticDir := "/static/"

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	return r
}

// Initiate web server
func main() {
	CompileStylesheets()

	router := Router()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9101",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server now live and listening at %s...", "127.0.0.1:9101")
	log.Fatal(srv.ListenAndServe())
}
