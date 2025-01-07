package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/pammalPrasanna/idi/internal/apps/idi"
	"github.com/pammalPrasanna/idi/internal/utils"
)

var ErrNoCommand = errors.New("missing command")

const none string = ""

func Main() error {
	projectName := flag.String("cp", none, "create project command: idi -cp [project name]")
	appName := flag.String("ca", none, "create app command: idi -ca [app name]")
	dbName := flag.String("cdb", none, "create db command: idi -cdb [mysql/postgres/sqlite3]")
	routerName := flag.String("cr", "httprouter", "create router command: idi -cr [chi/httprouter/mux]")
	isAuth := flag.Bool("auth", false, "create authentication flag: idi -cp -auth [default no auth]")
	isPaseto := flag.Bool("paseto", false, "create paseto flag: idi -cp -auth -paseto [default JWT]")
	showVersion := flag.Bool("v", false, "display version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", utils.Version())
		return nil
	}

	idi, err := idi.New(*projectName, *appName, *dbName, *routerName, *isAuth, *isPaseto)
	if err != nil {
		return err
	}
	err = idi.Create()
	if err != nil {
		return err
	}

	if idi != nil {
		return nil
	}
	return ErrNoCommand
}
