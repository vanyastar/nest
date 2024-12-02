package main

import "fmt"

type Default struct{}

func (c Default) Help() {
	fmt.Println(`
		Usage:

			nest <command> [arguments]

		The commands are:

			generate [method] [project directory] [name] 		Generate files for project in specified directory 
			- methods: <project|module|controller|service>
	`)
}
