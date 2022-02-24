package main

import (
	"log"
	"math/rand"
	"net/http"
)

const honeypotPort = ":7668"

func main() {
	http.HandleFunc("/", dumpBytes)
	err := http.ListenAndServe(honeypotPort, nil)
	log.Fatal("server stopped:" + err.Error())
}

func dumpBytes(w http.ResponseWriter, req *http.Request) {
	randomBytes := make([]byte, 10_000_000)
	w.Header().Add("Content-Type", "text/html")
	if _, err := rand.Read(randomBytes); err != nil {
		log.Fatal("failed generating bytes:" + err.Error())
	}
	if _, err := w.Write(randomBytes); err != nil {
		log.Fatal("failed to write bytes:" + err.Error())
	}
}
