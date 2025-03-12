package main

import (
	"fmt"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	server := http.Server{}
	server.Addr = ":8080"
	server.Handler = serveMux

	server.ListenAndServe()
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
