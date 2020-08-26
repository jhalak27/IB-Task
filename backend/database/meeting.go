package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Meeting struct {
	Pid          string `json:"pid"`
	Participant1 string `json:"participant1"`
	Participant2 string `json:"participant2"`
	Start        string `json:"start"`
	End          string `json:"end"`
}

//InsertMeeting to insert details of meeting into the database
func InsertMeeting(meeting Meeting, db *sql.DB) (int, error) {
	// db := connectDB()
	// defer db.Close()
	sqlStatement := `
	INSERT INTO meeting (participant1, participant2, start, end)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, meeting.Participant1, meeting.Participant2, meeting.Start, meeting.End).Scan(&id)
	if err != nil {
		panic(err)
	}
	// user.Id = id
	fmt.Println("New record ID is:", id)
	return id, err
}

//GetMeeting to get details of meeting from the database
func GetMeeting(pid int, db *sql.DB) (Meeting, error) {
	// db := connectDB()
	// defer db.Close()
	fmt.Println("starting query")
	var ans Meeting
	query := ` SELECT * FROM meeting WHERE pid = $1;`
	err := db.QueryRow(query, pid).Scan(&ans.Pid, &ans.Participant1, &ans.Participant2, &ans.Start, &ans.End)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	fmt.Println(ans)
	return ans, err
}

//GetAllMeeting to get details of all meeting from the database
func GetAllMeeting(db *sql.DB) ([]Meeting, error) {
	// db := connectDB()
	// defer db.Close()
	fmt.Println("starting query")
	var ans []Meeting
	query := ` SELECT * FROM meeting;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Unable to scan the row. %v", err)
		// handle this error better than this
		// panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var meet Meeting
		err = rows.Scan(&meet.Pid, &meet.Participant1, &meet.Participant2, &meet.Start, &meet.End)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		ans = append(ans, meet)
	}

	fmt.Println(ans)
	return ans, err
}

//UpdateMeeting to update details of meeting in the database
func UpdateMeeting(meeting Meeting, db *sql.DB) error {
	sqlStatement := `
	UPDATE meeting
	SET participant1 = $2, participant2 = $3, start = $4, end = $5
	WHERE pid = $1;`
	_, err := db.Exec(sqlStatement, meeting.Pid, meeting.Participant1, meeting.Participant2, meeting.Start, meeting.End)
	if err != nil {
		panic(err)
	}
	return err
}

//DeleteMeeting to delete details of meeting from the database
func DeleteMeeting(pid int, db *sql.DB) error {
	sqlStatement := `
	DELETE FROM meeting
	WHERE pid = $1;`
	_, err := db.Exec(sqlStatement, pid)
	if err != nil {
		panic(err)
	}
	return err
}
