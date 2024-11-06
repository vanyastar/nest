package nest

import "net/http"

func sseMidware() MidWare {
	return func(c *Ctx) {
		if c.Res().Flusher == nil {
			if f, ok := c.Res().ResponseWriter.(http.Flusher); ok {
				c.Res().Flusher = f
			} else {
				c.Error(http.StatusInternalServerError, "The flusher is not supported")
				return
			}
		}
		c.Res().Header().Set("Content-Type", "text/event-stream")
		c.Res().Header().Set("Cache-Control", "no-cache")
		c.Res().Header().Set("Connection", "keep-alive")
		c.Res().Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}
