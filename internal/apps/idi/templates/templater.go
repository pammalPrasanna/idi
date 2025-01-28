package templates

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pammalPrasanna/idi/internal/utils"
)

type Templater struct {
	AppNames    []string
	projectPath string
	ProjectName string
	DBName      string
	AppName     string
	RouterName  string
	Alias       string
	IsAuth      bool
	IsPaseto    bool
}

func New(projectPath, projectName, dbName, routerName, alias string, appNames []string, isAuth, isPaseto bool) *Templater {
	return &Templater{
		AppNames:    appNames,
		projectPath: projectPath,
		ProjectName: projectName,
		DBName:      dbName,
		RouterName:  routerName,
		Alias:       alias,
		IsAuth:      isAuth,
		IsPaseto:    isPaseto,
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

	// authentication
	if t.IsAuth {
		// create auth folder
		if err := t.createFolders(t.projectPath, authFolders); err != nil {
			return err
		}

		if t.IsPaseto {
			if err := t.createFiles(t.projectPath, pasetoFiles); err != nil {
				return err
			}
		} else {
			if err := t.createFiles(t.projectPath, jwtFiles); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t Templater) CreateApp() error {
	for _, appName := range t.AppNames {
		t.AppName = appName
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
	}

	return nil
}

func (t Templater) applyChoices(path string) string {
	if strings.Contains(path, "{") {
		path = strings.ReplaceAll(path, appKey, t.AppName)
		path = strings.ReplaceAll(path, routerKey, t.RouterName)
		path = strings.ReplaceAll(path, dbKey, t.DBName)
		path = strings.ReplaceAll(path, projectKey, t.ProjectName)
		path = strings.ReplaceAll(path, aliasKeys, t.Alias)
	}
	return path
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
