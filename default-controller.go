package nest

import (
	"strings"

	"github.com/vanyastar/nest/nestlog"
)

type DefaultControllerFunc func(c *DefaultController)

type DefaultController struct {
	appContextPath string
	*handlerManager
}

func (c *DefaultController) Sse() {
	c.Use(sseMidware())
}

func (c *DefaultController) Redirect(code int, location string) {
	c.Use(redirectMidWare(location, code))
}

func (c *DefaultController) Get(path string, ef *EndFunc) {
	c.apply("GET", path, ef)
}

func (c *DefaultController) Post(path string, ef *EndFunc) {
	c.apply("POST", path, ef)
}

func (c *DefaultController) Put(path string, ef *EndFunc) {
	c.apply("PUT", path, ef)
}

func (c *DefaultController) Delete(path string, ef *EndFunc) {
	c.apply("DELETE", path, ef)
}

func (c *DefaultController) Patch(path string, ef *EndFunc) {
	c.apply("PATCH", path, ef)
}

func (c *DefaultController) Head(path string, ef *EndFunc) {
	c.apply("HEAD", path, ef)
}

func (c *DefaultController) Options(path string, ef *EndFunc) {
	c.apply("OPTIONS", path, ef)
}

func (c *DefaultController) Connect(path string, ef *EndFunc) {
	c.apply("CONNECT", path, ef)
}

func (c *DefaultController) Trace(path string, ef *EndFunc) {
	c.apply("TRACE", path, ef)
}

func (c *DefaultController) apply(method, path string, ef *EndFunc) {
	c.setPath(strings.TrimSpace(method + " " + c.appContextPath + path))
	c.setEndpoint(ef)
	nestlog.Log("Router", "Mapped "+method+": "+c.appContextPath+path)
}

func newDefaultController(appContextPath string, handlerManager *handlerManager) *DefaultController {
	return &DefaultController{
		appContextPath: appContextPath,
		handlerManager: handlerManager,
	}
}
