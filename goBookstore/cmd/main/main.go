package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anuvratnehra/goProjectsFCC/goBookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegsiterBookStoreRoutes(r)
	http.Handle("/", r)
	err := http.ListenAndServe(":9010", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Running server on port 9010/n")
}
