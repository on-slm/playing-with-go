package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	res := "home"
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	http.Handle("/", r)
}
