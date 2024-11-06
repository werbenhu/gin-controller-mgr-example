package controller

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	globalRegistry *Registry
	once           sync.Once
)

// Registry - Global route registry to manage registered controllers
type Registry struct {
	mu     sync.RWMutex
	routes map[string]Controller
}

// Register - Registers a controller with a specified path
func (r *Registry) Register(path string, controller Controller) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.routes[path] = controller
}

// GetRoutes - Retrieves all registered routes
func (r *Registry) GetRoutes() map[string]Controller {
	r.mu.RLock()
	defer r.mu.RUnlock()
	routes := make(map[string]Controller)
	for k, v := range r.routes {
		routes[k] = v
	}
	return routes
}

// GetRegistry - Singleton function to get the global route registry
func GetRegistry() *Registry {
	once.Do(func() {
		globalRegistry = &Registry{
			routes: make(map[string]Controller),
		}
	})
	return globalRegistry
}

// Register - Helper function to register a controller in the global registry
func Register(path string, controller Controller) {
	GetRegistry().Register(path, controller)
}

// Controller - Defines the controller interface
type Controller interface {
	// Register - Registers a nested controller
	Register(path string, controller Controller)
	// Init - Initializes routes for the controller
	Init(router gin.IRouter)
}

// BaseController - Base controller implementation
type BaseController struct {
	router     gin.IRouter
	subRouters map[string]Controller
}

// NewBaseController - Creates a new base controller instance
func NewBaseController() *BaseController {
	return &BaseController{
		subRouters: make(map[string]Controller),
	}
}

// Register - Registers a sub-controller for a given path
func (b *BaseController) Register(path string, controller Controller) {
	b.subRouters[path] = controller
}

// Init - Initializes routes and sub-controllers
func (b *BaseController) Init(router gin.IRouter) {
	b.router = router
	// Initialize all sub-controllers
	for path, controller := range b.subRouters {
		controller.Init(router.Group(path))
	}
}

// ControllerManager - Manages controllers and routing with a Gin engine
type ControllerManager struct {
	engine     *gin.Engine
	controller Controller
}

// NewControllerManager - Creates a new controller manager instance
func NewControllerManager(engine *gin.Engine) *ControllerManager {
	return &ControllerManager{
		engine:     engine,
		controller: NewBaseController(),
	}
}

// Register - Registers a top-level route with a controller
func (m *ControllerManager) Register(path string, controller Controller) {
	m.controller.Register(path, controller)
}

// Init - Initializes all registered routes from the global registry
func (m *ControllerManager) Init() {
	// Retrieve all routes from the global registry
	routes := GetRegistry().GetRoutes()
	for path, controller := range routes {
		controller.Init(m.engine.Group(path))
	}
}
