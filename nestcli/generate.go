package main

import (
	"log"

	"github.com/vanyastar/nest/nestcli/tpl"
)

type Generate struct{}

func (c Generate) Project() error {

	dir, name := arguments.Pop(), arguments.Pop()

	if err := tpl.CreateProjectDirs(dir); err != nil {
		return err
	}
	if err := tpl.CreateModFile(dir, name); err != nil {
		return err
	}
	if err := tpl.CreateMainFile(dir, name); err != nil {
		return err
	}
	if err := tpl.CreateConfigFile(dir, name); err != nil {
		return err
	}
	if err := tpl.CreateAppFile(dir, name); err != nil {
		return err
	}
	if err := tpl.RunDependencies(dir); err != nil {
		return err
	}

	log.Print("- project created")
	return nil
}

func (c Generate) Module() error {
	dir, name := arguments.Pop(), arguments.Pop()

	if err := tpl.CreateModuleDirs(dir, name); err != nil {
		return err
	}

	if err := tpl.CreateControllerFile(dir+"/"+name, name); err != nil {
		return err
	}

	if err := tpl.CreateServiceFile(dir+"/"+name, name); err != nil {
		return err
	}

	log.Println("- module `" + name + "` created")
	return nil
}

func (c Generate) Controller() error {
	dir, name := arguments.Pop(), arguments.Pop()

	if err := tpl.CreateControllerFile(dir, name); err != nil {
		return err
	}

	log.Println("- controller `" + name + "` created")
	return nil
}

func (c Generate) Service() error {
	dir, name := arguments.Pop(), arguments.Pop()

	if err := tpl.CreateServiceFile(dir, name); err != nil {
		return err
	}

	log.Println("- service `" + name + "` created")
	return nil
}
