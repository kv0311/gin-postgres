package main

import (
	"fmt"
	"log"

	routes "github.com/cavdy-play/go_db/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("jdkjskdsj")
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}
