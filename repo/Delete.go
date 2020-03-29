package repo

import (
	_ "github.com/lib/pq"
)

func Delete(ID int) (err error) {
	db, err := connect()
	if err != nil {
		return
	}
	_, err = db.Query("delete from	staff.staff_card where id=$1", ID)
	if err != nil {
		return
	}
	return
}
