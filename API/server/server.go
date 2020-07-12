package server

import (
	"log"
	"net/http"
)

func Run(path string) {
	http.HandleFunc("/files/", getFiles(path))
	http.HandleFunc("/file/", getFile(path))
	http.HandleFunc("/", hello)
	log.Println("Server runnung on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
