package main

import "html/template"

func loadHTML(name ...string) (t *template.Template) {
	for i := range name {
		name[i] = "tmpl/" + name[i] + ".html"
	}
	t = template.Must(template.ParseFiles(name...))
	return
}
