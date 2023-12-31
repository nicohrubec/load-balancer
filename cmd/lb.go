package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRootLoadBalancer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "echo!\n")
	printRequest(r)
}

func main() {
	http.HandleFunc("/", getRootLoadBalancer)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
