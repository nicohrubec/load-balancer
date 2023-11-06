package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "echo!\n")
	printRequest(r)
}

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
