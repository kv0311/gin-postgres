package handler

import (
	"encoding/json"
	"fmt"
	"gin-demo/model"
	"gin-demo/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStudent(c *gin.Context) {
	type myRequest struct {
		ID int `json:"id"`
	}
	request := new(myRequest)
	err := c.Bind(request)
	if err != nil {
		fmt.Println(err)
	}
	//init redis
	client, err := Redis()
	if err != nil {
		fmt.Println(err)
	}
	//get data from redis
	cacheResult, err := client.Get("student=" + strconv.Itoa(request.ID)).Result()
	if err != nil {
		fmt.Errorf("error when get data from redis : %s", err)
	}
	student := model.Student{}
	// case has data in redis
	if cacheResult != "" {
		json.Unmarshal([]byte(cacheResult), &student)
		c.JSON(400, gin.H{
			"data": student,
		})
		//case not have data in redis
	} else {
		student, err := repo.Get(request.ID)
		if err != nil {
			return
		}
		marshalStudent, _ := json.Marshal(student)
		//after get data from postgresql set data to redis
		client.Set("student="+strconv.Itoa(request.ID), string(marshalStudent), 0)
		c.JSON(400, gin.H{
			"data": student,
		})
	}
}
