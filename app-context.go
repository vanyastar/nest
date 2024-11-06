package nest

import (
	logService "github.com/vanyastar/nest/log-service"
)

type AppContextFunc func(c *AppContext)

type AppContext struct {
	*handlerManager
	// For future use cases
	servers *httpServers
	router  *router
}

// Add the middleware globally, but it will still depend on where you are going to use it.
// At the global level or at the level of several controllers.
func (c *AppContext) UseGlobal(fn ...MidWare) {
	c.globalMidwares = append(c.globalMidwares, fn...)
}

// Static Server Controller
func (c *AppContext) Static(path string, dir string, fn DefaultControllerFunc) {
	logService.Log("StaticController", "Mapped "+path)
	c.Controller(path, fn, staticMidware(dir))
}

// Default Controller - Here you can override the default middlewares for this controller and assign your own.
func (c *AppContext) Controller(path string, fn DefaultControllerFunc, midWares ...MidWare) {
	if len(midWares) > 0 {
		c.Use(midWares...)
	} else {
		logService.Log("DefaultController", "Mapped "+path)
	}
	c.lockControllerMidwares()
	fn(newDefaultController(path, c.handlerManager))
	c.clearHandlerManager()
}

func newAppContext(r *router, s *httpServers) *AppContext {
	return &AppContext{
		servers:        s,
		router:         r,
		handlerManager: newHandlerManager(r),
	}
}
