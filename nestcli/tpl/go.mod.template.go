package tpl

import "fmt"

func ModTemplate(name string) []byte {
	return []byte(fmt.Sprintf(`module %s

go 1.23 // or newer
`, name))
}
