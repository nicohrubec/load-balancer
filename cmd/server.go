package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRootServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "echo from server!\n")
	printRequest(r)
}

func main() {
	http.HandleFunc("/", getRootServer)

	err := http.ListenAndServe(":10000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
