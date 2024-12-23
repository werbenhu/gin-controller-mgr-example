package controller

import "github.com/gin-gonic/gin"

// Initialization function to set up and register the controller hierarchy and routes
func init() {
	// Create the grandson controller and register it under the son controller at the "/grandson" path
	grandson := NewGrandsonController()
	Register("/api/father/son/grandson", grandson)

	son := NewSonController()
	Register("/api/father/son", son)

	// Create the father controller
	father := NewFatherController()
	Register("/", father)
}

// Example father controller
type FatherController struct {
	*BaseController
}

// Constructor function to create a new father controller instance
func NewFatherController() *FatherController {
	return &FatherController{
		BaseController: NewBaseController(),
	}
}

// Initialize routes for the father controller
func (u *FatherController) Init(ctx *Context, router gin.IRouter) {
	// Register specific route handlers for the father controller
	router.GET("/hello", u.Hello)    // Route for greeting
	router.POST("/create", u.Create) // Route for creating a new resource
}

// Handler function for the hello route in the father controller
func (u *FatherController) Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Father hello"})
}

// Handler function for creating a new resource in the father controller
func (u *FatherController) Create(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Father created"})
}

// Example son controller
type SonController struct {
	BaseController
}

// Constructor function to create a new son controller instance
func NewSonController() *SonController {
	return &SonController{}
}

// Initialize routes for the son controller
func (a *SonController) Init(ctx *Context, router gin.IRouter) {
	// Register specific route handlers for the son controller
	router.GET("/hello", a.Hello) // Route for greeting
	router.GET("/:id", a.Detail)  // Route for detailed view based on ID
}

// Handler function for the hello route in the son controller
func (a *SonController) Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Son Hello"})
}

// Handler function for displaying details based on ID in the son controller
func (a *SonController) Detail(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Son detail", "id": id})
}

// Example grandson controller
type GrandsonController struct {
	BaseController
}

// Constructor function to create a new grandson controller instance
func NewGrandsonController() *GrandsonController {
	return &GrandsonController{}
}

// Initialize routes for the grandson controller
func (a *GrandsonController) Init(ctx *Context, router gin.IRouter) {
	// Register specific route handlers for the grandson controller
	router.GET("/hello", a.Hello)   // Route for greeting
	router.GET("/detail", a.Detail) // Route for detailed view
}

// Handler function for the hello route in the grandson controller
func (a *GrandsonController) Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Grandson Hello"})
}

// Handler function for displaying detailed information in the grandson controller
func (a *GrandsonController) Detail(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Grandson detail", "id": id})
}
