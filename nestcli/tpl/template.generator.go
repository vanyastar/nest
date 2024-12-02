package tpl

import "os"

func CreateProjectDirs(dir string) error {
	if err := os.MkdirAll(dir+"/app", 0666); err != nil {
		return err
	}
	if err := os.MkdirAll(dir+"/configs", 0666); err != nil {
		return err
	}
	if err := os.MkdirAll(dir+"/common", 0666); err != nil {
		return err
	}
	return nil
}

func CreateModuleDirs(dir, name string) error {
	fDir := dir + "/" + name
	if err := os.MkdirAll(fDir, 0666); err != nil {
		return err
	}
	if err := os.MkdirAll(fDir+"/guards", 0666); err != nil {
		return err
	}
	if err := os.MkdirAll(fDir+"/dto", 0666); err != nil {
		return err
	}
	if err := os.MkdirAll(fDir+"/validators", 0666); err != nil {
		return err
	}
	if err := os.MkdirAll(fDir+"/interceptors", 0666); err != nil {
		return err
	}
	if err := os.MkdirAll(fDir+"/middlewares", 0666); err != nil {
		return err
	}
	return nil
}

func CreateControllerFile(dir, name string) error {
	return os.WriteFile(dir+"/"+name+".controller.go", ControllerTemplate(name), 0666)
}

func CreateServiceFile(dir, name string) error {
	return os.WriteFile(dir+"/"+name+".service.go", ServiceTemplate(name), 0666)
}

func CreateMainFile(dir, name string) error {
	return os.WriteFile(dir+"/main.go", MainTemplate(name), 0666)
}

func CreateModFile(dir, name string) error {
	return os.WriteFile(dir+"/go.mod", ModTemplate(name), 0666)
}

func CreateConfigFile(dir, name string) error {
	return os.WriteFile(dir+"/configs/http-server.config.go", CreateConfigTemplate(name), 0666)
}

func CreateAppFile(dir, name string) error {
	return os.WriteFile(dir+"/app/app.service.go", CreateAppTemplate(name), 0666)
}
