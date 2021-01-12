package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//
var db *gorm.DB
var err error
//
// task structure
type Task struct {
	gorm.Model
	id  uint 	`gorm:"primaryKey"`
	Name string `json:"name"`
	Done bool    `json:"done"`
}
//
// Main method for go
func main()  {
	//fmt.Println("start main task list \n")
	db, err = gorm.Open(sqlite.Open("task-list.sqlite"), &gorm.Config{})
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
    //reqBody, _ := ioutil.ReadAll(r.Body)
    //var task Task
	//json.Unmarshal(reqBody, &task)
	r.ParseForm()
	name := r.FormValue("name")
	//done := r.FormValue("done")
	task := Task{ Name: name }
    //fmt.Println("Endpoint Hit: Creating New Task ", task, " name ", name )
    db.Create(&task) 
    fmt.Println("Endpoint Hit: Creating New Task ", task )
    json.NewEncoder(w).Encode(task)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to HomePage!")
    fmt.Println("Endpoint Hit: HomePage")
}

func returnAllTasks(w http.ResponseWriter, r *http.Request){
	fmt.Println(" returnAllTasks")
	tasks := []Task{}
	db.Find(&tasks)
	fmt.Println("Endpoint Hit: returnAllTasks")
	json.NewEncoder(w).Encode(tasks)
}
//
func TaskHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println(" returnTask")
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println("id ", id)
	var task Task
	//tasks := []Task{}
	db.Find(&task, "id")
	//
	r.ParseForm()
	name := r.FormValue("name")
	task.Name = name
	db.Save(task)
	//
	json.NewEncoder(w).Encode(task)
}
   
func TaskDeleteHandler( w http.ResponseWriter, r *http.Request ){
	fmt.Println(" TaskDeleteHandler")
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println("id ", id)
	//tasks := []Task{}
	db.Delete(&Task{}, id)
	//db.Delete()
	json.NewEncoder(w).Encode("{'response': 'success'}")
}

//
func TaskUpdateHandler( w http.ResponseWriter, r *http.Request ){
	fmt.Println(" TaskUpdateHandler")
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println("id ", id)
	var task Task
	db.Find(&task, "id")
	json.NewEncoder(w).Encode(task)
}
//
// handle function
func handleRequests(){
    log.Println("Starting development server at http://127.0.0.1:8000/")
    log.Println("Quit the server with CONTROL-C.")
    // creates a new instance of a mux router
    r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/create", createNewTask).Methods("POST")
	r.HandleFunc("/all", returnAllTasks )
	r.HandleFunc("/health", HealthCheckHandler )
	r.HandleFunc("/task/{id:[0-9]+}", TaskHandler)
	r.HandleFunc("/task-delete/{id:[0-9]+}", TaskDeleteHandler) // need to be use the same routing but differente methods like DELETE
	r.HandleFunc("/task-update/{id:[0-9]+}", TaskUpdateHandler) // need to be use the same routing but differente methods like PUT
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

