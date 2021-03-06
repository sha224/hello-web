package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type helloWebHandler struct {
	port     string
	hostname string
	name     string
}

type contextKey struct {
	key string
}

var connContextKey = &contextKey{"conn"}

func getLocalAddress(r *http.Request) (string, bool) {
	conn, ok := r.Context().Value(connContextKey).(net.Conn)
	if ok {
		return conn.LocalAddr().String(), true
	}
	return "", false
}

func (handler helloWebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received on port %v from %v\n", handler.port, r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, World! You have reached port %v of %v from %v.\n", handler.port, handler.name, r.RemoteAddr)
	fmt.Fprintf(w, "\nHost Details:\n")
	fmt.Fprintf(w, "Name: %v\n", handler.name)
	fmt.Fprintf(w, "Hostname: %v\n", handler.hostname)

	addr, ok := getLocalAddress(r)
	if ok {
		fmt.Fprintf(w, "Server Address: %v\n", addr)
	}

	addrs, err := net.InterfaceAddrs()
	if err == nil {
		fmt.Fprintf(w, "All Addresses: %v\n", addrs)
	}

	fmt.Fprintf(w, "\nRequest Details:\n")
	fmt.Fprintf(w, "Method: %v\n", r.Method)
	fmt.Fprintf(w, "URL: %v\n", r.URL)
	fmt.Fprintf(w, "Protocol: %v\n", r.Proto)
	fmt.Fprintf(w, "Host: %v\n", r.Host)
	fmt.Fprintf(w, "Header: %v\n", r.Header)
	fmt.Fprintf(w, "Body: %v\n", r.Body)
	fmt.Fprintf(w, "Content Length: %v\n", r.ContentLength)
	fmt.Fprintf(w, "Transfer Encoding: %v\n", r.TransferEncoding)
	fmt.Fprintf(w, "Form Data: %v\n", r.Form)
	fmt.Fprintf(w, "Remote Address: %v\n", r.RemoteAddr)
}

func main() {
	hostname, _ := os.Hostname()

	name := os.Getenv("NAME")
	if name == "" {
		name = hostname
	}

	ports := os.Getenv("PORTS")
	if ports == "" {
		ports = "8080"
	}

	portStrings := strings.Split(ports, ",")

	for _, port := range portStrings {
		log.Printf("Listening on port %v", port)

		server := http.Server{
			Addr:    ":" + port,
			Handler: helloWebHandler{port, hostname, name},
			ConnContext: func(ctx context.Context, c net.Conn) context.Context {
				return context.WithValue(ctx, connContextKey, c)
			},
		}
		go server.ListenAndServe()
	}

	for {
	}
}
