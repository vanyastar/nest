package nest

import (
	"net/http"
	"os"
)

type StaticServerConfig struct {
	RootDir   string
	MimeTypes map[string]string
}

func staticMidware(dir string) MidWare {
	return func(c *Ctx) {
		file, err := os.Open(dir + c.Req().URL.Path)
		if err == nil {
			defer file.Close()
			info, err := file.Stat()
			if err != nil {
				c.Error(http.StatusInternalServerError, err.Error())
				return
			}
			if !info.IsDir() {
				http.ServeContent(c.Res(), c.Req().Request, info.Name(), info.ModTime(), file)
				return
			}
		}
		c.Next()
	}
}
