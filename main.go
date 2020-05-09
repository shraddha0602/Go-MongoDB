package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shraddha0602/golang-mongodb/config"
	"github.com/shraddha0602/golang-mongodb/routes"
)

func main() {

	//connect db
	config.Connect()

	//initialize router
	router := gin.Default()

	routes.Routes(router)

	//Listen to the router with port
	log.Fatal(router.Run("localhost:8800"))
}
