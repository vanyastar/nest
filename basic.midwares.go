package nest

import "net/http"

func headerMidWare(key, value string) MidWare {
	return func(c *Ctx) {
		c.Res().Header().Set(key, value)
		c.Next()
	}
}

func maxBodySizeMidWare(maxBodySize int64) MidWare {
	return func(c *Ctx) {
		c.Req().Body = http.MaxBytesReader(c.Res(), c.Req().Body, maxBodySize)
		c.Next()
	}
}

func redirectMidWare(location string, code int) MidWare {
	return func(c *Ctx) {
		c.Res().Header().Set("Location", location)
		c.Next()
		c.Res().WriteHeader(code)
	}
}
