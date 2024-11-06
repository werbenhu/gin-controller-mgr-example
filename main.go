package main

import (
	"gotest/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// Create a controller manager
	manager := controller.NewControllerManager(engine)

	// Initialize all routes
	manager.Init()

	// Start the server
	engine.Run(":8080")
}
