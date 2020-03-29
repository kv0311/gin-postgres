package repo

import (
	_ "github.com/lib/pq"
)

func Update(Name string, ID int) (err error) {
	db, err := connect()
	if err != nil {
		return
	}
	//Check connect with db success or not
	_, err = db.Query("update staff.staff_card set name = $1 where id = $2", Name, ID)
	// Query Update db
	if err != nil {
		return
	}
	//Check Update db fail or success
	return
}
