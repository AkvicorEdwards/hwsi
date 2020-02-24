package index

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"hwsi/config"
	"hwsi/theme/ori"
	"net/http"
)

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	var err error
	if config.Data.Path.Theme == "ori" {
		t, err = template.New("ori").Parse(ori.Index)
	}else {
		t, err = template.ParseFiles(config.Data.Path.Theme + "index.html")
	}

	if err != nil {
		log.Error(err)
		return
	}

	if err = t.Execute(w, map[string]interface{}{"title":config.Data.Server.Title}); err != nil {
		log.Error(err)
		_, _ = fmt.Fprintf(w, "%v", "Error")
	}
	return
}

// Index page
func FileIndex(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/index", http.FileServer(http.Dir(config.Data.Path.Work))).ServeHTTP(w, r)
}