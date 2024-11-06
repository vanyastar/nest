package app

import (
	"example/app/cats"
	"example/app/dogs"
	staticServer "example/app/static-server"

	nest "github.com/vanyastar/nest"
)

func App(c *nest.AppContext) {
	// if needed global level of middlewares
	c.UseGlobal()

	// Start cats and dogs controller
	cats.NewCatsController(c)
	dogs.NewDogsController(c)

	// Start static controller
	staticServer.NewStaticController(c)
}
