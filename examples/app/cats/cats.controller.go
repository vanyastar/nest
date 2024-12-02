package cats

import (
	catsDtos "example/app/cats/dtos"
	"net/http"

	"github.com/vanyastar/nest"
)

type catsController struct {
	catsService *CatsService

	getNameString             nest.EndFunc
	getNameMap                nest.EndFunc
	setNameStruct             nest.EndFunc
	setNameStructNoValidation nest.EndFunc
	setNameMap                nest.EndFunc
	sessionExample            nest.EndFunc
	pushNotification          nest.EndFunc
	jsonError                 nest.EndFunc
}

// Here is your logic for that controller
func NewCatsController(c *nest.AppContext) {
	this := &catsController{
		catsService: GetCatsService,
	}

	c.Header("Server", "Example1")
	c.Controller("/api/cats", func(c *nest.DefaultController) {

		c.Header("Server", "Controller itself")
		c.Get("", &this.getNameString)
		this.getNameString = func(c *nest.Ctx) error {
			return c.SendString("Hello, this is an API controller.")
		}

		c.Header("Server", "Method1")
		c.Get("/name", &this.getNameMap)
		this.getNameMap = func(c *nest.Ctx) error {
			return c.Send(
				map[string]string{
					"Name": this.catsService.GetName(),
				},
			)
		}

		// Post request with fully automated DTO struct
		c.Post("/name-struct", &this.setNameStruct)
		this.setNameStruct = func(c *nest.Ctx) error {
			dto := new(catsDtos.SetCatName)
			if err := c.DtoParser(dto); err != nil {
				return c.Error(http.StatusBadRequest, err.Error())
			}
			this.catsService.SetName(dto.Name)
			return c.Send(
				map[string]string{
					"Name": this.catsService.GetName(),
				},
			)
		}

		c.Post("/name-struct-no-validation", &this.setNameStructNoValidation)
		this.setNameStructNoValidation = func(c *nest.Ctx) error {
			dto := new(catsDtos.SetCatNameNoValidation)
			if err := c.BodyParser(dto); err != nil {
				return c.Error(http.StatusBadRequest, err.Error())
			}
			if err := this.catsService.SetName(dto.Name); err != nil {
				return c.Error(http.StatusBadRequest, err.Error())
			}
			return c.Send(
				map[string]string{
					"Name": this.catsService.GetName(),
				},
			)
		}

		c.Post("/name-map", &this.setNameMap)
		this.setNameMap = func(c *nest.Ctx) error {
			dto := make(map[string]any)
			if err := c.BodyParser(&dto); err != nil {
				return c.Error(http.StatusBadRequest, err.Error())
			}
			if data, ok := dto["name"].(string); ok {
				if err := this.catsService.SetName(data); err != nil {
					return c.Error(http.StatusBadRequest, err.Error())
				}
			} else {
				return c.Error(http.StatusBadRequest, "Assert to string error")
			}
			return c.Send(
				map[string]string{
					"Name": this.catsService.GetName(),
				},
			)
		}

		// Session example
		c.Get("/session", &this.sessionExample)
		this.sessionExample = func(c *nest.Ctx) error {
			session := c.Session()
			if value, ok := session.GetValue("CatName"); ok {
				return c.SendString(value.(string))

			}
			session.SetValue("CatName", this.catsService.GetName()).Save(c)
			return c.Error(http.StatusNotFound, "Setting up a session and the cat's name in it. Refresh this page")
		}

		c.Get("/json-error", &this.jsonError)
		this.jsonError = func(c *nest.Ctx) error {
			err := map[string]string{
				"message":   "simple custom json error",
				"otherData": "12335",
			}
			return c.Error(500, err)
		}

		// Server sent event example
		c.Sse()
		c.Get("/sse", &this.pushNotification)
		this.pushNotification = func(c *nest.Ctx) error {
			return this.catsService.pushNotification(c)
		}
	})
}
