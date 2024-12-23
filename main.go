package main

import (
	"gotest/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// Create a controller manager
	manager := controller.NewControllerManager(engine)

	ctx := &controller.Context{}

	// Initialize all routes
	manager.Init(ctx)

	// Start the server
	engine.Run(":8080")
}
