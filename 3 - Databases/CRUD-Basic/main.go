package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("CRUD")

	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreatUser).Methods(http.MethodPost)

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
