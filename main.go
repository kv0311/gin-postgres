package main

import (
	"gin-demo/routes"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	// Route Handlers / Endpoints
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=khanhvinh dbname=postgres password=khanhvinh1998")
	// defer db.Close()
	// type Person struct {
	// 	ID          int    `gorm:"id"`
	// 	Name        string `gorm:"name"`
	// 	School      string `gorm:"school"`
	// 	Class       string `gorm:"class"`
	// 	Description string `gorm:"description"`
	// }
	// user := Person{}
	// db.First(&user)
	// fmt.Println(db)
	// fmt.Println(err)
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}
