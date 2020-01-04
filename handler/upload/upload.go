package upload

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"hwsi/config"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if t, err := template.ParseFiles(config.Data.Path.Theme+"upload.html"); err != nil {
			log.Error(err)
			return
		} else if err := t.Execute(w, map[string]interface{}{"Title":config.Data.Server.Title}); err != nil {
			log.Error(err)
			return
		}
	} else {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			log.Error(err)
			return
		}
		file, handler, err := r.FormFile("filename")
		if err != nil {
			_, _ = fmt.Fprintf(w, "%v", "Upload error")
			log.Info(err)
			return
		}
		ext := filepath.Ext(handler.Filename)
		if !checkFileType(ext) {
			_, _ = fmt.Fprintf(w, "%v", "Illegal file type")
			return
		}
		filename := strconv.FormatInt(time.Now().Unix(), 10) + "_" + handler.Filename
		f, _ := os.OpenFile(config.Data.Path.Work + filename, os.O_CREATE | os.O_WRONLY, 0660)
		if _, err = io.Copy(f, file); err != nil {
			_, _ = fmt.Fprintf(w, "%v", "Upload failed")
			return
		}
		fileDir, _ := filepath.Abs(config.Data.Path.Work + filename)
		_, _ = fmt.Fprintf(w, "%v", filename+" Upload successful: "+fileDir)

	}
}

// File types not allowed to upload
func checkFileType(name string) bool {
	ext := []string{".exec"}
	for _, v := range ext {
		if v == name {
			return false
		}
	}
	return true
}
