package templates

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pammalPrasanna/idi/internal/utils"
)

type Templater struct {
	projectPath string
	ProjectName string
	DBName      string
	AppName     string
	RouterName  string
}

func New(projectPath, projectName, dbName, appName, routerName string) *Templater {
	return &Templater{
		projectPath: projectPath,
		ProjectName: projectName,
		DBName:      dbName,
		AppName:     appName,
		RouterName:  routerName,
	}
}

func (t Templater) CreateProjectFolder() error {
	if err := os.Mkdir(t.projectPath, perm); err != nil {
		return err
	}
	return nil
}

func (t Templater) CreateFramework() error {
	// create framework folders
	if err := t.createFolders(t.projectPath, frameworkFolders); err != nil {
		return err
	}

	if t.DBName != none {
		// if yes generate db infra
		if err := t.createFiles(t.projectPath, dbFiles); err != nil {
			return err
		}
	}

	// create framework default files
	if err := t.createFiles(t.projectPath, frameworkDefaultFiles); err != nil {
		return err
	}

	return nil
}

func (t Templater) CreateApp() error {
	// create app default folders
	if err := t.createFolders(t.projectPath, appDefaultFolders); err != nil {
		return err
	}

	// create app default files
	if err := t.createFiles(t.projectPath, appDefaultFiles); err != nil {
		return err
	}

	// create db folders and files if specified
	if t.DBName != "" {
		if err := t.createFolders(t.projectPath, appDBFolders); err != nil {
			return err
		}

		if err := t.createFiles(t.projectPath, appDBFiles); err != nil {
			return err
		}
	}

	return nil
}

func (t Templater) applyChoices(path string) string {
	tempPath := strings.ReplaceAll(path, appKey, t.AppName)
	tempPath = strings.ReplaceAll(tempPath, routerKey, t.RouterName)
	tempPath = strings.ReplaceAll(tempPath, dbKey, t.DBName)
	tempPath = strings.ReplaceAll(tempPath, projectKey, t.ProjectName)
	return tempPath
}

func (t Templater) createFiles(path string, filesMap map[string]*itemplate) error {
	for _, itemplate := range filesMap {
		tempPath := t.applyChoices(itemplate.path)
		filePath := filepath.Join(path, tempPath)

		fileTemplate := template.Must(template.New("").Funcs(templateFunctions).Parse(string(itemplate.content)))

		if err := utils.CreateWriteTemplate(filePath, fileTemplate, t); err != nil {
			return err
		}
	}
	return nil
}

func (t Templater) createFolders(path string, foldersMap map[string]string) error {
	for _, v := range foldersMap {
		tempPath := t.applyChoices(v)
		if err := os.MkdirAll(filepath.Join(path, tempPath), perm); err != nil {
			return err
		}
	}
	return nil
}
