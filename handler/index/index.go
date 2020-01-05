package index

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"hwsi/config"
	"net/http"
)

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(config.Data.Path.Theme + "index.html")
	if err := t.Execute(w, map[string]interface{}{"Title":config.Data.Server.Title}); err != nil {
		log.Error(err)
		_, _ = fmt.Fprintf(w, "%v", "Error")
	}
	return
}

// Index page
func FileIndex(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/index", http.FileServer(http.Dir(config.Data.Path.Work))).ServeHTTP(w, r)
}