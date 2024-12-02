package tpl

func CreateAppTemplate(name string) []byte {
	return []byte(`package app

import (
	"github.com/vanyastar/nest"
)

func App(c *nest.AppContext) {
	c.UseGlobal()
}
`)
}
