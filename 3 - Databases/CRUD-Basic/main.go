package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("CRUD")

	router := mux.NewRouter()

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
