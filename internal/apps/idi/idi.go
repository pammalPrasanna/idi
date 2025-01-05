package idi

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"github.com/pammalPrasanna/idi/internal/apps/idi/templates"
	"github.com/pammalPrasanna/idi/internal/utils"
)

type Idi struct {
	none        string
	projectName string
	appName     string
	dbName      string
	projectPath string
	routerName  string
}

var (
	dbList     = [...]string{"mysql", "postgres", "sqlite3"}
	routerList = [...]string{"chi", "httprouter", "mux"}
)

func New(projectName, appName, dbName, routerName string) (*Idi, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if err = validateDBName(dbName); err != nil {
		return nil, err
	}
	if err = validateRouterName(routerName); err != nil {
		return nil, err
	}

	return &Idi{
		none:        "",
		projectName: projectName,
		appName:     appName,
		dbName:      dbName,
		routerName:  routerName,
		projectPath: filepath.Join(cwd, projectName),
	}, nil
}

func (i Idi) Create() error {
	if i.projectName != i.none {
		// validate project name
		// generate project
		t := templates.New(i.projectPath, i.projectName, i.dbName, i.appName, i.routerName)
		if err := t.CreateProjectFolder(); err != nil {
			return err
		}

		if err := t.CreateFramework(); err != nil {
			return err
		}

		if err := utils.InitGoMod(i.projectName, i.projectPath); err != nil {
			return err
		}
	}

	if i.appName != i.none {
		// ensure -ca command is executed from idi project folder
		if i.projectName == i.none {
			fmt.Println("project name empty")
			prjDir, err := i.idiProjectExists()
			if err != nil {
				return err
			}
			i.projectPath = prjDir
		}

		// OR
		// -ca command is executed along with -cp command
		// validate app name already doesn't exists

		// if cwd + our apps apth exists --> project exists else no project found
		t := templates.New(i.projectPath, i.projectName, i.dbName, i.appName, i.routerName)
		if err := t.CreateApp(); err != nil {
			return err
		}
	}

	return nil
}

func validateDBName(dbName string) error {
	// check given db is present in dbmap
	if dbName != "" && !slices.Contains(dbList[:], dbName) {
		return errors.New("db not found: " + dbName)
	}
	return nil
}

func validateRouterName(routerName string) error {
	// check given db is present in dbmap
	if routerName != "" && !slices.Contains(routerList[:], routerName) {
		return errors.New("db not found: " + routerName)
	}
	return nil
}

func (i Idi) idiProjectExists() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	goMOD := filepath.Join(cwd, "go.mod")

	if _, err := os.Stat(goMOD); errors.Is(err, fs.ErrNotExist) {
		return "", errors.New("go project not found")
	}

	appsDir := filepath.Join(cwd, "/internal/apps")

	if _, err := os.Stat(appsDir); errors.Is(err, fs.ErrNotExist) {
		return "", errors.New("idi project structure not found")
	}

	return cwd, nil
}
