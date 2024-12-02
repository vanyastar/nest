package tpl

import (
	"strings"
)

// MainTemplate generates a Go main template with the provided name
func MainTemplate(name string) []byte {
	// Define the template with placeholders for the dynamic parts
	template := `package main

import (
	"{{name}}/app"
	"{{name}}/configs"
	"os"
	"os/signal"
	"syscall"

	"github.com/vanyastar/nest"
)

func main() {
	nestApp := nest.CreateApp(app.App,
		configs.NewHttpServer(),
		configs.NewHttp2Server(),
	)
	nestApp.ListenAndServe()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	nestApp.Shutdown()
}
`

	// Create a replacer that will replace all occurrences of {{name}} with the provided name
	replacer := strings.NewReplacer("{{name}}", name)

	// Apply the replacer to the template
	result := replacer.Replace(template)

	// Return the final template as a byte slice
	return []byte(result)
}
