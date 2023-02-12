package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		environment := os.Getenv("Environment")
		w.Write([]byte("hello devops v12: " + environment))
	})
	log.Print("Server started, Listening on port 10000...")
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
