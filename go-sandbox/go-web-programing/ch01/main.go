package main

import (
	"log"
	"net/http"
	"sync"
	"text/template"
	"path/filepath"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

// ServerHTTPはHTTPリクエストを処理する
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(`
	// 		<html>
	// 			<head>
	// 				<title>チャット</title>
	// 			</head>
	// 			<body>
	// 				チャットしましょう
	// 			</body>
	// 		</html>
	// 	`))
	// })
	// Webサーバーを始める
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}