package main

import (
	"fmt"
	"net/http"

	"github.com/Parzivalxx/go-course/pkg/handlers"
)

const portNumber = "localhost:8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
