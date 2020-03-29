package handler

import (
	"gin-demo/model"
	"gin-demo/repo"

	"github.com/gin-gonic/gin"
)

func AddStudent(c *gin.Context) {
	type myRequest struct {
		ID          int    `json:"id" validate:"required"`
		Name        string `json:"name" validate:"required"`
		School      string `json:"school" validate:"required"`
		Class       string `json:"class" validate:"required"`
		Description string `json:"description"`
	}
	request := new(myRequest)
	err := c.Bind(request)
	//bind request params to request variable

	if err != nil {
		c.AbortWithError(400, err)
	}
	//check err of bind function
	student := model.Student{}
	err = repo.Add(request.ID, request.Name, request.School, request.Class, request.Description)
	//call repo.Add to query database
	if err != nil {
		c.AbortWithError(400, err)
	}
	//check err when query database
	c.JSON(400, gin.H{
		"status":  200,
		"message": "Thêm học sinh thành công",
		"data":    student,
	})
}
