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
		fmt.Println("Run: ")
		fmt.Printf("1. cd %s \n", *projectName)
		fmt.Printf("2. go fmt .%s...\n", string(os.PathSeparator))
		fmt.Println("3. go mod tidy")
		fmt.Printf("4. go run .%s...", string(os.PathSeparator))
		return nil
	}
	return ErrNoCommand
}
