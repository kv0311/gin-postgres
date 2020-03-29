package handler

import (
	"gin-demo/repo"

	"github.com/gin-gonic/gin"
)

func UpdateName(c *gin.Context) {
	type myRequest struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	request := new(myRequest)
	err := c.Bind(request)
	//bind request params to request variable
	if err != nil {
		c.AbortWithError(400, err)
	}
	//check err of bind function
	err = repo.Update(request.Name, request.ID)
	if err != nil {
		c.AbortWithError(400, err)
	}
	//check err when query database
	c.JSON(400, gin.H{
		"status":  200,
		"message": "Update tên học sinh thành công",
	})
}
