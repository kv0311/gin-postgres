package repo

import (
	"log"

	_ "github.com/lib/pq"
)

func Add(ID int, Name, Class, School, Description string) (err error) {

	db, err := connect()
	if err != nil {
		return
	}
	_, err = db.Query("insert into	staff.staff_card values ($1,$2,$3,$4,$5)", ID, Name, Class, School, Description)
	if err != nil {
		log.Fatal(err)
	}
	return
}
