package tpl

import "strings"

func ServiceTemplate(name string) []byte {
	// Define the template with placeholders
	template := `package {{name}}

type {{uName}}Service struct {
	name string
}

var Get{{uName}}Service = &{{uName}}Service{}

func (this *{{uName}}Service) GetName() string {
	return ""
}
`
	toUpperName := strings.ToUpper(name[:1]) + name[1:]

	// Apply the replacer to the template
	result := strings.ReplaceAll(template, "{{name}}", name)
	result = strings.ReplaceAll(result, "{{uName}}", toUpperName)

	return []byte(result)
}
