//go:build go1.16

package generator

import (
	"embed"
	"text/template"
)

//go:embed iface.tmpl
var content embed.FS

func addEmbeddedTemplates(tmpl *template.Template) *template.Template {
	return template.Must(tmpl.ParseFS(content, "*.tmpl"))
}
