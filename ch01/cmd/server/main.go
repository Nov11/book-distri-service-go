package main

import (
	"github.com/Nov11/proglog/ch01/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", handler.NewHTTPServer())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
