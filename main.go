package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

//task structure
type Task struct {
	gorm.Model
	id  uint
	name string
	done bool
}
//
// Main method for go
func main()  {
	fmt.Print("start main task list")
	db, err := gorm.Open(sqlite.Open("task-list.db"), &gorm.Config{})
	//
	fmt.Print("err = ", err)
	if err != nil {
		fmt.Print("failed to connect")
		panic("failed to connect database")
	}
	//
	// Migrate the schema
	db.AutoMigrate(&Task{})
	//
	task := Task{id: 1, name: "test", done: false }
	//
	log.Print("task = ", task)
	// Create
	result := db.Create(&task)
	fmt.Print("result.RowsAffected " , result)
	//
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
	//
}