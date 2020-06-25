package handler

import (
	"hwsi/config"
	"hwsi/handler/index"
	"hwsi/handler/upload"
	"net/http"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

// All Prefix
func Init() {
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = index.HomePage
	mux["/upload"] = upload.Upload
	mux["/index"] = index.FileIndex
}

// Prefix Handler
type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	http.StripPrefix("/", http.FileServer(http.Dir(config.WorkDir))).ServeHTTP(w, r)
}

