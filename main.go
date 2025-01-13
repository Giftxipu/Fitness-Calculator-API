package main

import (
  	"github.com/gin-gonic/gin"
  	"fitness-calculator-api/handlers"
)

func main(){
  router := gin.Default()

	// Register the fitness handler
	router.POST("/fitness", handlers.FitnessHandler)

	// Start the server on port 8080
	router.Run(":8080")
}
