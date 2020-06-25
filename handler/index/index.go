package index

import (
	"fmt"
	"hwsi/config"
	"hwsi/tpl"
	"log"
	"net/http"
)

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	if err := tpl.Index.Execute(w, map[string]interface{}{"title":config.Title}); err != nil {
		log.Println(err)
		_, _ = fmt.Fprintf(w, "%v", "Error")
	}
	return
}

// Index page
func FileIndex(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/index", http.FileServer(http.Dir(config.WorkDir))).ServeHTTP(w, r)
}