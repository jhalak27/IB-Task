package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	Db "./database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// type Participant struct {
// 	Id       string `json:"id"`
// 	Name     string `json:"name"`
//  Email    string `json:"email"`
// 	Slot    string `json:"slot"``
// }

// type Meeting struct {
// 	Pid          string `json:"pid"`
// 	Participant1 string `json:"participant1"`
// 	Participant2 string `json:"participant2"`
// 	Start        string `json:"start"`
// 	End          string `json:"end"`
// }
var Participants []Db.Participant
var Meetings []Db.Meeting
var db *sql.DB

func connectDB() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "588800"
		dbname   = "task"
	)
	fmt.Println("Connecting to the database")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createMeeting(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var meeting Db.Meeting
	json.Unmarshal(reqBody, &meeting)
	// fmt.Println(len(Users))
	// user.Id = strconv.Itoa(len(Users) + 1)
	Meetings = append(Meetings, meeting)
	x, _ := Db.InsertMeeting(meeting, db)
	meeting.Pid = strconv.Itoa(x)
	json.NewEncoder(w).Encode(meeting)
}

func listAllMeetings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUser")
	Meetings, err := Db.GetAllMeeting(db)
	if err != nil {
		log.Fatalf("Unable to get all meetings. %v", err)
	}
	json.NewEncoder(w).Encode(Meetings)
}

func editMeeting(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var meeting Db.Meeting
	json.Unmarshal(reqBody, &meeting)
	err := Db.UpdateMeeting(meeting, db)
	if err != nil {
		fmt.Println(err)
	}
}

func deleteMeeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	pid, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	err = Db.DeleteMeeting(pid, db)
	if err != nil {
		fmt.Println(err)
	}
}

func listAllParticipants(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUser")
	Participants, err := Db.GetAllParticipant(db)
	if err != nil {
		log.Fatalf("Unable to get all participants. %v", err)
	}
	json.NewEncoder(w).Encode(Participants)
}

func editParticipant(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var participant Db.Participant
	json.Unmarshal(reqBody, &participant)
	err := Db.UpdateParticipant(participant, db)
	if err != nil {
		fmt.Println(err)
	}
}

func handleRequest() {
	fmt.Println("Routers")
	myRouter := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/participants", listAllParticipants).Methods("GET")
	myRouter.HandleFunc("/participant/{id}", editParticipant).Methods("POST")
	myRouter.HandleFunc("/meetings", listAllMeetings).Methods("GET")
	myRouter.HandleFunc("/meeting", createMeeting).Methods("POST")
	myRouter.HandleFunc("/meeting/{id}", editMeeting).Methods("POST")
	myRouter.HandleFunc("/meeting/{id}", deleteMeeting).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(myRouter)))
}

func main() {
	db = connectDB()
	defer db.Close()
	handleRequest()
}
