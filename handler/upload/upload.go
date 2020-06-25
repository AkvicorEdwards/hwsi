package upload

import (
	"fmt"
	"html/template"
	"hwsi/config"
	"hwsi/theme"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Upload file
func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var t *template.Template
		var err error

		t, err = template.New("ori").Parse(theme.Upload)

		if err != nil {
			log.Println(err)
			return
		}

		if err := t.Execute(w, map[string]interface{}{"title":config.Title}); err != nil {
			log.Println(err)
			return
		}
	} else {

		if err := r.ParseMultipartForm(32 << 20); err != nil {
			log.Println(err)
			return
		}
		if r.FormValue("password") != config.Password {
			_, _ = fmt.Fprintf(w, "%v", "Wrong password")
			return
		}
		file, handler, err := r.FormFile("filename")
		if err != nil {
			_, _ = fmt.Fprintf(w, "%v", "Upload error")
			log.Println(err)
			return
		}
		filename := handler.Filename
		f, _ := os.OpenFile(config.UploadDir + filename, os.O_CREATE | os.O_WRONLY, 0660)
		if _, err = io.Copy(f, file); err != nil {
			_, _ = fmt.Fprintf(w, "%v", "Upload failed")
			return
		}
		fileDir, _ := filepath.Abs(config.UploadDir + filename)
		_, _ = fmt.Fprintf(w, "%v", filename+" Upload successful: "+fileDir)

	}
}
