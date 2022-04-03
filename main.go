package main

import (
	"encoding/base64"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

const honeypotPort = ":3001"

func main() {
	http.HandleFunc("/", dumpBytes)
	err := http.ListenAndServe(honeypotPort, nil)
	log.Fatal("server stopped:" + err.Error())
}

func dumpBytes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var randomBytes []byte
	if strings.Contains(r.Host, "localhost") {
		randomBytes = make([]byte, 10_000)
	} else {
		randomBytes = make([]byte, 10_000_000)
	}
	if _, err := rand.Read(randomBytes); err != nil {
		log.Fatal("failed generating bytes:" + err.Error())
	}
	b64 := base64.NewEncoder(base64.RawStdEncoding, w)
	if err := b64.Close(); err != nil {
		log.Println("failed to close b64 encoder", err.Error())
	}

	w.Header().Add("Content-Type", "text/html")
	htmlHeader := []byte(`
	<!DOCTYPE html>
	<html>
		<head><title>this is a honeypot. welcome</title></head>
		<body>
		<style>
			.b64 {
				width: 100%;
				word-wrap: break-word;
			}
		</style>
			<div class="b64">
			`)
	htmlFooter := []byte(`
			</div>
		</body>
	</html>`)
	if _, err := w.Write(htmlHeader); err != nil {
		log.Println("failed writing header for page:", err.Error())
	}
	if _, err := b64.Write(randomBytes); err != nil {
		log.Println("failed encoding bytes to base64", err.Error())
	}
	if _, err := w.Write(htmlFooter); err != nil {
		log.Println("failed writing footer for page:", err.Error())
	}
}
