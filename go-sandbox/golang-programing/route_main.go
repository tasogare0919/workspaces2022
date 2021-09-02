package main

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}
	templates := template.Must(template.PerseFiles(files...))
	threads, err := data.Threads()
	if err == nill {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
