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
	"io/ioutil"
	"encoding/json"
)
//
var db *gorm.DB
var err error
//
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
	//fmt.Println("start main task list \n")
	db, err := gorm.Open(sqlite.Open("task-list.sqlite"), &gorm.Config{})
	//
	//mt.Print("\n db = ", db)
	if err != nil {
		fmt.Println("failed to connect")
		panic("failed to connect database")
	}
	//
	// Migrate the schema
	db.AutoMigrate(&Task{})
	//
	//
	handleRequests()
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

func createNewTask(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
    reqBody, _ := ioutil.ReadAll(r.Body)
    var task Task
    json.Unmarshal(reqBody, &task)
    db.Create(&task) 
    fmt.Println("Endpoint Hit: Creating New Task")
    json.NewEncoder(w).Encode(task)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to HomePage!")
    fmt.Println("Endpoint Hit: HomePage")
}

func returnAllTasks(w http.ResponseWriter, r *http.Request){
	tasks := []Task{}
	db.Find(&tasks)
	fmt.Println("Endpoint Hit: returnAllTasks")
	json.NewEncoder(w).Encode(tasks)
}
   

func handleRequests(){
    log.Println("Starting development server at http://127.0.0.1:8000/")
    log.Println("Quit the server with CONTROL-C.")
    // creates a new instance of a mux router
    r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/create", createNewTask)
	r.HandleFunc("/all", returnAllTasks )
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
}


/*
func CreateTaskHandler(w http.ResponseWriter, r *http.Request){

    // A very simple health check.
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	// need to be reformated the code
	task := Task{ name: "test", done: false }
	//
	fmt.Print("\n task = ", task)
	// Create
	result := db.Create(&task)
	fmt.Print("\n result.RowsAffected " , result)
	//
    // In the future we could report back on the status of our DB, or our cache
    // (e.g. Redis) by performing a simple PING, and include them in the response.
    io.WriteString(w, `{"task": created }`)
}*/