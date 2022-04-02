package main

import (
	"WebGo/handler"
	"log"
	"net/http"
)

func main() {
	about := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ini Halaman About"))
	}

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/helo", handler.Hello)
	http.HandleFunc("/mario", handler.Mario)
	http.HandleFunc("/about", about)
	http.HandleFunc("/portofolio", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ini Halaman Portofolio"))
	})
	http.HandleFunc("/post-get", handler.PostGet)
	http.HandleFunc("/form", handler.Form)
	http.HandleFunc("/proses", handler.Proses)

	fileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	http.HandleFunc("/product", handler.Product)
	log.Println("starting web server at http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
