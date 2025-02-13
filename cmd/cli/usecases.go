package cli

import "errors"

// InterruptCase 1
// we cannot create db files on already existing application
// we interrupt the CLI

func InterruptCase1(dbName, appNames, projectName string) error {
	switch {
	case dbName == "":
		return nil
	case dbName != "":
		if appNames == "" {
			return errors.New("cannot scaffold DB without create app command '-ca'")
		}
	}
	return nil
}
