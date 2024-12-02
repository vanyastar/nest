package app

import (
	"example/app/cats"
	"example/app/dogs"
	"example/app/static"

	"github.com/vanyastar/nest"
)

func App(c *nest.AppContext) {
	// if needed global level of middlewares
	c.UseGlobal()

	// Start cats and dogs controller
	cats.NewCatsController(c)
	dogs.NewDogsController(c)

	// Start static controller
	static.NewStaticController(c)
}
