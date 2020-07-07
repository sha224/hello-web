package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type helloWebHandler struct {
	port     string
	hostname string
	name     string
}

func (handler helloWebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received on port %v from %v\n", handler.port, r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, World! You have reached port %v of %v from %v.\n", handler.port, handler.name, r.RemoteAddr)
	fmt.Fprintf(w, "\nHost Details:\n")
	fmt.Fprintf(w, "Name: %v\n", handler.name)
	fmt.Fprintf(w, "Hostname: %v\n", handler.hostname)
	fmt.Fprintf(w, "\nRequest Details:\n")
	fmt.Fprintf(w, "Method: %v\n", r.Method)
	fmt.Fprintf(w, "URL: %v\n", r.URL)
	fmt.Fprintf(w, "Protocol: %v\n", r.Proto)
	fmt.Fprintf(w, "Header: %v\n", r.Header)
	fmt.Fprintf(w, "Body: %v\n", r.Body)
	fmt.Fprintf(w, "Content Length: %v\n", r.ContentLength)
	fmt.Fprintf(w, "Transfer Encoding: %v\n", r.TransferEncoding)
	fmt.Fprintf(w, "Host: %v\n", r.Host)
	fmt.Fprintf(w, "Form Data: %v\n", r.Form)
	fmt.Fprintf(w, "Remote Address: %v\n", r.RemoteAddr)
	fmt.Fprintf(w, "Request URI: %v\n", r.RequestURI)
}

func main() {
	hostname, _ := os.Hostname()

	name := os.Getenv("NAME")
	if name == "" {
		name = hostname
	}

	ports := os.Getenv("PORTS")
	if ports == "" {
		ports = "80"
	}

	portStrings := strings.Split(ports, ",")

	for _, port := range portStrings {
		log.Printf("Listening on port %v", port)
		go http.ListenAndServe(":"+port, helloWebHandler{port, hostname, name})
	}

	for {
	}
}