package tpl

import (
	"fmt"
	"html/template"
)

var err error
var errCnt = 0

var Index *template.Template
var Upload *template.Template

func Parse() {
	Index = add("index", index)
	Upload = add("upload", upload)
	fmt.Printf("Parsing the html template was completed with %d errors\n", errCnt)
}

func add(name, tpl string) (t *template.Template) {
	t, err = template.New(name).Parse(tpl)
	if err != nil {
		errCnt++
		fmt.Println(err)
	}
	return
}
