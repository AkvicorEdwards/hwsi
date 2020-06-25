package index

import (
	"fmt"
	"html/template"
	"hwsi/config"
	"hwsi/theme"
	"log"
	"net/http"
)

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	var err error
	t, err = template.New("ori").Parse(theme.Index)
	if err != nil {
		log.Println(err)
		return
	}

	if err = t.Execute(w, map[string]interface{}{"title":config.Title}); err != nil {
		log.Println(err)
		_, _ = fmt.Fprintf(w, "%v", "Error")
	}
	return
}

// Index page
func FileIndex(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/index", http.FileServer(http.Dir(config.WorkDir))).ServeHTTP(w, r)
}