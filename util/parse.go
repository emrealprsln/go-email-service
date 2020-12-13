package util

import (
	"bytes"
	"fmt"
	"html/template"
)

const (
	templateBasePath = "templates"
)

func ParseTemplate(file string, data interface{}) string {
	buffer := new(bytes.Buffer)
	t, _ := template.ParseFiles(fmt.Sprintf("%s/%s", templateBasePath, file))
	t.Execute(buffer, data)
	return buffer.String()
}
