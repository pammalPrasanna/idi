package idi

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/pammalPrasanna/idi/internal/apps/idi/templates"
	"github.com/pammalPrasanna/idi/internal/utils"
)

type Idi struct {
	appNames    []string
	none        string
	projectName string
	dbName      string
	projectPath string
	routerName  string
	alias       string
	isAuth      bool
	isPaseto    bool
}

var (
	dbList     = [...]string{"mysql", "postgres", "sqlite3"}
	routerList = [...]string{"chi", "httprouter", "mux"}
)

func New(projectName, appNames, dbName, routerName, alias string, isAuth, isPaseto bool) (*Idi, error) {
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

	appNamesArr := getAppNames(appNames)

	return &Idi{
		projectName: projectName,
		appNames:    appNamesArr,
		dbName:      dbName,
		routerName:  routerName,
		alias:       alias,
		projectPath: filepath.Join(cwd, projectName),
		none:        "",
		isAuth:      isAuth,
		isPaseto:    isPaseto,
	}, nil
}

func (i Idi) Create() error {
	if i.projectName != i.none {
		// validate project name
		// generate project
		t := templates.New(i.projectPath, i.projectName, i.dbName, i.routerName, i.alias, i.appNames, i.isAuth, i.isPaseto)
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

	if len(i.appNames) != 0 {
		// ensure -ca command is executed from idi project folder
		if i.projectName == i.none {
			prjDir, err := i.idiProjectExists()
			if err != nil {
				return err
			}
			i.projectPath = prjDir
			dirs := strings.Split(i.projectPath, string(os.PathSeparator))
			i.projectName = dirs[len(dirs)-1]
		}

		// OR
		// -ca command is executed along with -cp command
		// validate app name already doesn't exists

		// if cwd + our apps apth exists --> project exists else no project found
		t := templates.New(i.projectPath, i.projectName, i.dbName, i.routerName, i.alias, i.appNames, i.isAuth, i.isPaseto)
		if err := t.CreateApp(); err != nil {
			return err
		}
	}

	return nil
}

func getAppNames(appNames string) []string {
	split := strings.Split(appNames, ",")
	for i, s := range split {
		split[i] = strings.TrimSpace(s)
	}
	return split
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

	appsDir := filepath.Join(cwd, "/internal/apps")

	_, err = os.Stat(appsDir)
	if err != nil {
		switch {
		case errors.Is(err, fs.ErrNotExist):
			return "", errors.New(`idi project structure not found:
				1. '{current_working_directory}/internal/apps'
				2. '{current_working_directory}/internal/dtos'`)

		default:
			return "", err
		}
	}

	return cwd, nil
}
