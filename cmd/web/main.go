package main

import (
	"fmt"
	"net/http"

	"github.com/irisida/goserver/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Application started on port: %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
