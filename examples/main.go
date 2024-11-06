package main

import (
	"example/app"
	"example/configs"
	"os"
	"os/signal"
	"syscall"

	nest "github.com/vanyastar/nest"
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
