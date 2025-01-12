package main

import (
	"BookMangement/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("We're running 🤪 ")
	log.Fatal(http.ListenAndServe(":8080", r))
}
