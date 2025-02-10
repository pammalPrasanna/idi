package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"without-alias/cmd/api"
	"without-alias/internal/lib"
)

func main() {
	idi, err := lib.Idi()
	if err != nil {
		fmt.Println(err)
		trace := string(debug.Stack())
		fmt.Println(trace)
		os.Exit(1)
	}

	
	// handle DB close on application exit
	defer func() {
		if err := idi.Postgres().Close(); err != nil {
			idi.Logger().Error("unable to close Postgres", "Postgres.Close()", err)
		} else {
			idi.Logger().Info("closed Postgres successfully")
		}
	}()
	

	err = api.Main(idi)
	if err != nil {
		trace := string(debug.Stack())
		idi.Logger().Error("error: ", err)
		idi.Logger().Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

