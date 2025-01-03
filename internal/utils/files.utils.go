package utils

import (
	"os"
	"os/exec"
	"text/template"
)

func CreateWriteTemplate(filePath string, templ *template.Template, t any) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = templ.Execute(file, t)
	if err != nil {
		return err
	}

	return nil
}

// InitGoMod initializes go.mod with the given project name
// in the selected directory
func InitGoMod(projectName string, projectDir string) error {
	if err := ExecuteCmd("go",
		[]string{"mod", "init", projectName},
		projectDir); err != nil {
		return err
	}

	return nil
}

// ExecuteCmd provides a shorthand way to run a shell command
func ExecuteCmd(command string, args []string, dir string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
