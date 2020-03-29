package handler

import (
	"gin-demo/repo"

	"github.com/gin-gonic/gin"
)

func DeleteStudent(c *gin.Context) {
	type myRequest struct {
		ID int `json:"id"`
	}
	request := new(myRequest)
	err := c.Bind(request)
	//bind request params to request variable
	if err != nil {
		c.AbortWithError(400, err)
	}
	//check err of bind function
	err = repo.Delete(request.ID)
	if err != nil {
		c.AbortWithError(400, err)
	}
	//check err when query database
	c.JSON(400, gin.H{
		"status":  200,
		"message": "Xóa học sinh thành công",
	})
}
