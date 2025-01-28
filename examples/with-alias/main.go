package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"with-alias/cmd/api"
	"with-alias/internal/lib"
)

func main() {
	rootApp, err := lib.RootApp()
	if err != nil {
		fmt.Println(err)
		trace := string(debug.Stack())
		fmt.Println(trace)
		os.Exit(1)
	}

	// handle DB close on application exit
	defer func() {
		if err := rootApp.Sqlite3().Close(); err != nil {
			rootApp.Logger().Error("unable to close Sqlite3", "Sqlite3.Close()", err)
		} else {
			rootApp.Logger().Info("closed Sqlite3 successfully")
		}
	}()

	err = api.Main(rootApp)
	if err != nil {
		trace := string(debug.Stack())
		rootApp.Logger().Error("error: ", err)
		rootApp.Logger().Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
