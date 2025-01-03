package templates

import (
	_ "embed"
)

// app files start

//go:embed files/app/dtos.go.tmpl
var dtoFile []byte

//go:embed files/app/initapp.go.tmpl
var initFile []byte

//go:embed files/app/domain.go.tmpl
var domainFile []byte

//go:embed files/app/irepo.go.tmpl
var irepoFile []byte

//go:embed files/app/repo.go.tmpl
var repoFile []byte

//go:embed files/app/application.go.tmpl
var appFile []byte

//go:embed files/app/controller.go.tmpl
var ctrlrFile []byte

//go:embed files/app/routes.go.tmpl
var routeFile []byte

// app files end

// framework files start

//go:embed files/framework/main.go.tmpl
var mainFile []byte

//go:embed files/framework/api.go.tmpl
var apiFile []byte

//go:embed files/framework/helpers.go.tmpl
var helpersFile []byte

//go:embed files/framework/rootapp.go.tmpl
var rootAppFile []byte

//go:embed files/framework/errors.go.tmpl
var errorsFile []byte

//go:embed files/framework/json.go.tmpl
var jsonFile []byte

//go:embed files/framework/server.go.tmpl
var serverFile []byte

//go:embed files/framework/logger.go.tmpl
var loggerFile []byte

//go:embed files/framework/types.go.tmpl
var typesFile []byte

//go:embed files/framework/db.go.tmpl
var dbFile []byte

//go:embed files/framework/validator.go.tmpl
var validatorFile []byte

//go:embed files/framework/rest.go.tmpl
var restFile []byte

//go:embed files/framework/auth.go.tmpl
var authFile []byte

//go:embed files/framework/metrics.go.tmpl
var metricsFile []byte

// framework files end
