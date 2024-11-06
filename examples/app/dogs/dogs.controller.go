package dogs

import (
	"example/app/cats"

	nest "github.com/vanyastar/nest"
)

type dogsController struct {
	dogsService *DogsService
	catsService *cats.CatsService

	getCatName          nest.EndFunc
	getCatNameThrowCats nest.EndFunc
}

func NewDogsController(c *nest.AppContext) {
	var this = &dogsController{
		dogsService: GetDogsService,
		catsService: cats.GetCatsService,
	}

	c.Controller("/api/dogs", func(c *nest.DefaultController) {
		// Get cat name throw dogs service
		c.Get("/cat-name", &this.getCatName)
		this.getCatName = func(c *nest.Ctx) error {
			return c.SendString(this.dogsService.GetCatName())
		}

		c.Get("/cat-name-from-cats-service", &this.getCatNameThrowCats)
		this.getCatNameThrowCats = func(c *nest.Ctx) error {
			// Get the name of the cat, through the cat service in the dog controller, instead of using catsService throw dogsService
			return c.SendString(this.catsService.GetName())
		}
	})
}
