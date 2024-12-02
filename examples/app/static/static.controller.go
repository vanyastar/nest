package static

import (
	"net/http"

	"github.com/vanyastar/nest"
)

type staticController struct {
	IndexFile nest.EndFunc
}

func NewStaticController(c *nest.AppContext) {
	this := &staticController{}

	c.Static("/", "./public", func(c *nest.DefaultController) {

		c.Get("", &this.IndexFile)
		this.IndexFile = func(c *nest.Ctx) error {
			if c.Req().URL.Path == "/" {
				return c.SendFile("./public/index.html")

			}
			return c.Error(http.StatusNotFound, "404 Not found")
		}

	})
}
