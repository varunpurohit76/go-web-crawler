package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/varunpurohit76/crawler/base"
	"github.com/varunpurohit76/crawler/graph"
	"github.com/varunpurohit76/crawler/scrapper"
	"github.com/varunpurohit76/crawler/sitemap"
)

func main() {
	base.InitConfig()
	if err := base.ConnectDb(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	base.InitLog()
	sitemap.Service.Init(scrapper.PageUrlExtract, graph.JsonView)

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
	r.HandleFunc("/scrap", sitemap.SitemapHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8765", r))
}
