package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes(router)

	fmt.Println("server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
