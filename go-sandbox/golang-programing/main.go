package main

import (
	"net/http"
)

func manin() {
	mux := http.NewServeMux()                                   //マルチプレクサ生成
	files := http.FileServer(http.Dir("/public"))               //ファイル返送
	mux.Handle("/static/", http.StripPrefix("/static/", files)) //リクエストURLからプレフィックスを削除する

	mux.HandleFunc("/", index) //ルートURLをハンドラ関数にリダイレクトする
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
