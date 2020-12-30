package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main()  {
	fmt.Print("start main task list")
	
	r := mux.NewRouter()
    //r.HandleFunc("/", HomeHandler)
    //r.HandleFunc("/products", ProductsHandler)
    //r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	//
}