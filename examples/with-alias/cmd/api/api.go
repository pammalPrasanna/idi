package api

import (
	 
	 "with-alias/internal/apps/users"
	 
	"with-alias/internal/lib"
)

func Main(rootApp lib.IApp) error {
	
	users.InitApp(rootApp)
	
	return rootApp.ServeHTTP()
}
