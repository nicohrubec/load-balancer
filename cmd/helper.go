package main

import (
	"fmt"
	"net/http"
	"strings"
)

func printRequest(r *http.Request) {
	addressParts := strings.Split(r.RemoteAddr, ":")
	fmt.Println("Received request from: ", addressParts[0])

	fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, r.Proto)
	fmt.Printf("Host: %s\n", r.Host)

	for key, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
