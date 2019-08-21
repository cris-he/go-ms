package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"go-gateway/models"
	"go-gateway/routes"
)

func main() {
	var router = mux.NewRouter()

	host := os.Getenv("HOST")
	if host == "" {
		host = "http://localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" //localhost
	}

	endpoint := models.EndPoint{host: host, port: port}

	routes.HandleRoutes(router, endpoint)

	fmt.Println("Ready to Go " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
