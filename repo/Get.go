package repo

import (
	"fmt"
	"gin-demo/model"

	_ "github.com/lib/pq"
)

func Get(ID int) (student model.Student, err error) {
	db, err := connect()
	if err != nil {
		return
	}
	rows, err := db.Query("select * from staff.staff_card where id=$1", ID)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		student := model.Student{}
		err = rows.Scan(&student.ID, &student.Name, &student.School, &student.Class, &student.Description)
		fmt.Println(student, err)
		return student, err
	}

	return
}
