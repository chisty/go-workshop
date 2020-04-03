package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello from the Server")
	})
	http.ListenAndServe(":9090", nil)

	fmt.Println("Server started ...")
}
