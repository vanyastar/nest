package tpl

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vanyastar/nest/nestlog"
)

func RunDependencies(dir string) error {
	nestlog.Log("CMD", "Installing dependencies")
	cDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Define the command you want to run and the directory to run it in
	cmd := exec.Command("go", "get", "github.com/vanyastar/nest@latest")
	cmd.Dir = filepath.Join(cDir, dir)

	// Set up the command's output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err = cmd.Run()
	if err != nil {
		return err
	}
	nestlog.Log("CMD", "Dependencies installed")
	return nil
}
