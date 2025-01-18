package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"asdf/cmd/api"
	"asdf/internal/lib"
)

func main() {
	idi, err := lib.Idi()
	if err != nil {
		fmt.Println(err)
		trace := string(debug.Stack())
		fmt.Println(trace)
		os.Exit(1)
	}
	defer idi.Sqlite3().Close()

	err = api.Main(idi)
	if err != nil {
		trace := string(debug.Stack())
		idi.Logger().Error("error: ", err)
		idi.Logger().Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
