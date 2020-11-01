package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err)
	}
}
