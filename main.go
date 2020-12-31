package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
    "io"
)

// task structure
type Task struct {
	gorm.Model
	id  uint `gorm:"primaryKey"`
	name string
	done bool
}
//
// Main method for go
func main()  {
	fmt.Print("start main task list \n")
	db, err := gorm.Open(sqlite.Open("task-list.sqlite"), &gorm.Config{})
	//
	fmt.Print("\n db = ", db)
	if err != nil {
		fmt.Print("failed to connect")
		panic("failed to connect database")
	}
	//
	// Migrate the schema
	db.AutoMigrate(&Task{})
	//
	//
	r := mux.NewRouter()
    //r.HandleFunc("/", HomeHandler)
    //r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	r.HandleFunc("/health", HealthCheckHandler )
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
//
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    // A very simple health check.
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // In the future we could report back on the status of our DB, or our cache
    // (e.g. Redis) by performing a simple PING, and include them in the response.
    io.WriteString(w, `{"alive": true}`)
}
//
func CreateTaskHandler(w http.ResponseWriter, r *http.Request){

    // A very simple health check.
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	// need to be reformated the code
	//task := Task{id: 1, name: "test", done: false }
	//
	//fmt.Print("\n task = ", task)
	// Create
	//result := db.Create(&task)
	//fmt.Print("\n result.RowsAffected " , result)
	//
    // In the future we could report back on the status of our DB, or our cache
    // (e.g. Redis) by performing a simple PING, and include them in the response.
    io.WriteString(w, `{"task": created }`)
}