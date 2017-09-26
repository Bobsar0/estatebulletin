package http

import "html/template"
import "strings"

var (
	indextmpl  = newAppTemplate("index.html")
	signuptmpl = newAppTemplate("signup.html")
)

type appTemplate struct {
	t *template.Template
}

func newAppTemplate(filename string) *appTemplate {
	path := strings.Join([]string{"templates", filename}, "/")
	tmpl := template.Must(template.ParseFiles("templates/base.html", path))
	return &appTemplate{
		t: tmpl,
	}
}
