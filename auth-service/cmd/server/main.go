package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"auth-service/gen/auth/v1/authv1connect" // generated by protoc-gen-connect-go
	"auth-service/handlers"
)

func main() {
	// Set up authentication server
	authServer := &handlers.AuthServer{}

	// Set up HTTP multiplexer
	mux := http.NewServeMux()

	// Set up authentication service handler
	path, handler := authv1connect.NewAuthServiceHandler(authServer)
	mux.Handle(path, handler)

	// Configure server address
	serverAddr := ":8080"

	// Use h2c to serve HTTP/2 without TLS
	server := &http.Server{
		Addr:    serverAddr,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	// Start the server
	fmt.Println("Server started at", serverAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: ", err)
	}
}
