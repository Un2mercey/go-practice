package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("My dream App"))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start HTTP Server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
