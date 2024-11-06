package nest

import "reflect"

type MidWare func(*Ctx)
type ControllerHandler func(*Ctx, *handler)
type EndFunc func(*Ctx) error

type handler struct {
	path     string
	midWares []MidWare
	endPoint *EndFunc
}

type handlerManager struct {
	*router
	*handler
	globalMidwares         []MidWare
	controllerMidwares     []MidWare
	methodMidwares         []MidWare
	controllerMidWaresLock bool
}

func (c *handlerManager) applyDefaultHandler() {
	if (len(c.path)) > 0 {
		tHandlers := c.globalMidwares
		tHandlers = append(tHandlers, c.controllerMidwares...)
		tHandlers = append(tHandlers, c.methodMidwares...)

		uniqMap := make(map[uintptr]byte)
		var filteredMidWares []MidWare

		for i := len(tHandlers) - 1; i >= 0; i-- {
			midWare := tHandlers[i]
			p := reflect.ValueOf(midWare).Pointer()
			if uniqMap[p] != 1 {
				filteredMidWares = append([]MidWare{midWare}, filteredMidWares...)
			}
			uniqMap[p] = 1
		}
		c.handler.midWares = filteredMidWares
		c.router.defaultHandler(c.handler)
		c.methodMidwares = c.methodMidwares[:0]
		c.handler = &handler{}
	}
}

func (c *handlerManager) setPath(path string) {
	c.path = path
}

func (c *handlerManager) Header(key, value string) {
	c.Use(headerMidWare(key, value))
}

func (c *handlerManager) MaxBodySize(maxBodySize int64) {
	c.Use(maxBodySizeMidWare(maxBodySize))
}

// Add middlewares for each level of application
func (c *handlerManager) Use(fn ...MidWare) {
	if c.controllerMidWaresLock {
		c.methodMidwares = append(c.methodMidwares, fn...)
	} else {
		c.controllerMidwares = append(c.controllerMidwares, fn...)
	}
}

func (c *handlerManager) setEndpoint(fn *EndFunc) {
	c.endPoint = fn
	c.applyDefaultHandler()
}

func (c *handlerManager) lockControllerMidwares() {
	c.controllerMidWaresLock = true
}

func (c *handlerManager) clearHandlerManager() {
	c.path = ""
	c.controllerMidWaresLock = false
	c.controllerMidwares = c.controllerMidwares[:0]
	c.methodMidwares = c.methodMidwares[:0]
	c.endPoint = nil
}

func newHandlerManager(router *router) *handlerManager {
	return &handlerManager{
		router:  router,
		handler: &handler{},
	}
}
