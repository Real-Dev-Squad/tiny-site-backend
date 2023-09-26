package main

import (
	"net/http"
	"url-shortener/routes"
)

func main() {
	http.HandleFunc("/", routes.MainHandler)
	http.ListenAndServe(":8080", nil)
}
