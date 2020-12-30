package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"log"
)

func main()  {
	fmt.Print("start main task list")
	
	r := mux.NewRouter()
    //r.HandleFunc("/", HomeHandler)
    //r.HandleFunc("/products", ProductsHandler)
    //r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	//
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}