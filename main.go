package main

import (
	"fmt"
	"grpc-middleware/middleware"
	"grpc-middleware/service/contacts"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	server := contacts.NewServer()
	server.Run()

	router := mux.NewRouter()
	router.Use(middleware.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Print(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}
