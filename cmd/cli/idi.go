package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/pammalPrasanna/idi/internal/apps/idi"
	"github.com/pammalPrasanna/idi/internal/utils"
)

var ErrNoCommand = errors.New("missing command")

const none string = ""

func Main() error {
	projectName := flag.String("cp", none, "create project with flag: idi -cp [project name]")
	appNames := flag.String("ca", none, "add one or more apps with flag: idi -ca [appname1,appname2]")
	dbName := flag.String("cdb", none, "add db with flag: idi -cdb [mysql/postgres/sqlite3]")
	routerName := flag.String("cr", "httprouter", "add router with flag: idi -cr [chi/httprouter/mux] (currently 'httprouter' only)")
	isAuth := flag.Bool("auth", false, "add JWT authentication with flag: idi -cp -auth")
	isPaseto := flag.Bool("paseto", false, "add Paseto instead of JWT with flag: idi -cp -auth -paseto")
	alias := flag.String("a", "idi", "change root app name with flag: idi -alias 'myRootApp'")
	showVersion := flag.Bool("v", false, "display version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("version: %s\n", utils.Version())
		return nil
	}

	idi, err := idi.New(*projectName, *appNames, *dbName, *routerName, *alias, *isAuth, *isPaseto)
	if err != nil {
		return err
	}
	err = idi.Create()
	if err != nil {
		return err
	}

	if idi != nil {
		fmt.Println("Run: ")
		fmt.Printf("1. cd %s \n", *projectName)
		fmt.Printf("2. go fmt .%s...\n", string(os.PathSeparator))
		fmt.Println("3. go mod tidy")
		fmt.Printf("4. go run .%s...", string(os.PathSeparator))
		return nil
	}
	return ErrNoCommand
}
