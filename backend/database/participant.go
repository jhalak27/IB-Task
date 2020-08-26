package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Participant struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Slot  string `json:"slot"`
}

//InsertParticipant to insert details of participant into the database
func InsertParticipant(participant Participant, db *sql.DB) (int, error) {
	// db := connectDB()
	// defer db.Close()
	sqlStatement := `
	INSERT INTO participant (name, email, slot)
	VALUES ($1, $2, $3)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, participant.Name, participant.Email, participant.Slot).Scan(&id)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	// participant.Id = id
	fmt.Println("New record ID is:", id)
	return id, err
}

//GetParticipant to get details of participant from the database
func GetParticipant(id int, db *sql.DB) (Participant, error) {
	// db := connectDB()
	// defer db.Close()
	fmt.Println("starting query")
	var ans Participant
	query := ` SELECT * FROM participant WHERE id = $1;`
	err := db.QueryRow(query, id).Scan(&ans.Id, &ans.Name, &ans.Email, &ans.Slot)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	fmt.Println(ans)
	return ans, err
}

//GetAllParticipant to get details of all participant from the database
func GetAllParticipant(db *sql.DB) ([]Participant, error) {
	// db := connectDB()
	// defer db.Close()
	fmt.Println("starting query")
	var ans []Participant
	query := ` SELECT * FROM participant;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Unable to scan the row. %v", err)
		// handle this error better than this
		// panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var par Participant
		err = rows.Scan(&par.Id, &par.Name, &par.Email, &par.Slot)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		ans = append(ans, par)

	}

	fmt.Println(ans)
	return ans, err
}

//UpdateParticipant to update details of participant in the database
func UpdateParticipant(participant Participant, db *sql.DB) error {
	sqlStatement := `
	UPDATE participant
	SET name = $2, email = $3, slot = $4
	WHERE id = $1;`
	_, err := db.Exec(sqlStatement, participant.Id, participant.Name, participant.Email, participant.Slot)
	if err != nil {
		panic(err)
	}
	return err
}

//DeleteParticipant to delate details of participant from the database
func DeleteParticipant(id int, db *sql.DB) error {
	sqlStatement := `
	DELETE FROM participant
	WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	return err
}
