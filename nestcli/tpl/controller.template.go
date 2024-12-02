package tpl

import (
	"strings"
)

// ControllerTemplate generates a Go controller template with the provided name
func ControllerTemplate(name string) []byte {
	// Define the template with placeholders for the dynamic parts
	template := `package {{name}}

import (
	"github.com/vanyastar/nest"
)

type {{name}}Controller struct {
	{{name}}Service *{{uName}}Service

	{{name}}Method nest.EndFunc
}

func New{{name}}Controller(c *nest.AppContext) {
	this := &{{name}}Controller{
		{{name}}Service: Get{{uName}}Service,
	}

	c.Controller("/", func(c *nest.DefaultController) {

		c.Get("", &this.{{name}}Method)
		this.{{name}}Method = func(c *nest.Ctx) error {
			return c.SendString(this.{{name}}Service.GetName())
		}
	})
}
`

	toUpperName := strings.ToUpper(name[:1]) + name[1:]

	// Apply the replacer to the template
	result := strings.ReplaceAll(template, "{{name}}", name)
	result = strings.ReplaceAll(result, "{{uName}}", toUpperName)

	// Return the final template as a byte slice
	return []byte(result)
}
