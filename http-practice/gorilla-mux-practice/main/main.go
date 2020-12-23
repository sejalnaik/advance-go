package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/products", productsHandler)
	r.HandleFunc("/articles/{category}", articlesHandler)

	//http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home!!!!!!!!!!!!!!!!!!!!!!!!!</h1>"))

}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Products!!!!!!!!!!!!!!!!!!!!!!!!!</h1>"))

}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
	w.Write([]byte("<h1>Articles!!!!!!!!!!!!!!!!!!!!!!!!!</h1>"))

}
