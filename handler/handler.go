package handler

import (
	"hwsi/config"
	myindex "hwsi/handler/index"
	"hwsi/handler/upload"
	"net/http"
	"regexp"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

// All Prefix
func Init() {
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = myindex.HomePage
	mux["/upload"] = upload.Upload
	mux["/index"] = myindex.FileIndex
}

// Prefix Handler
type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	if ok, _ := regexp.MatchString("/css/", r.URL.String()); ok {
		http.StripPrefix("/css/", http.FileServer(http.Dir(config.Data.Path.Theme+"css/"))).ServeHTTP(w, r)
	} else if ok, _ := regexp.MatchString("/js/", r.URL.String()); ok {
		http.StripPrefix("/js/", http.FileServer(http.Dir(config.Data.Path.Theme+"js/"))).ServeHTTP(w, r)
	} else if ok, _ := regexp.MatchString("/img/", r.URL.String()); ok {
		http.StripPrefix("/img/", http.FileServer(http.Dir(config.Data.Path.Theme+"img/"))).ServeHTTP(w, r)
	} else {
		http.StripPrefix("/", http.FileServer(http.Dir(config.Data.Path.Work))).ServeHTTP(w, r)
	}
}

