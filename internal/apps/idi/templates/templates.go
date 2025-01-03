package templates

type itemplate struct {
	path    string
	content []byte
}

var frameworkFolders = map[string]string{
	"cmd":            "cmd/{project_name}/api",
	"apps":           "internal/apps",
	"dtos":           "internal/dtos",
	"framework":      "internal/idi",
	"infrastructure": "internal/infrastructure",
}

var frameworkDefaultFiles = map[string]*itemplate{
	// cmd files start
	"api.go": {
		path:    "cmd/{project_name}/api/api.go",
		content: apiFile,
	},
	"config.go": {
		path:    "cmd/{project_name}/api/helpers.go",
		content: helpersFile,
	},
	"root.app.go": {
		path:    "cmd/{project_name}/api/{project_name}.root.go",
		content: rootAppFile,
	},
	"main.go": {
		path:    "cmd/{project_name}/main.go",
		content: mainFile,
	},
	// cmd files ends

	// framework files start
	"errors.go": {
		path:    "internal/idi/errors.go",
		content: errorsFile,
	},
	"json.go": {
		path:    "internal/idi/json.go",
		content: jsonFile,
	},
	"server.go": {
		path:    "internal/idi/server.go",
		content: serverFile,
	},
	"logger.go": {
		path:    "internal/idi/slogger.go",
		content: loggerFile,
	},
	"types.go": {
		path:    "internal/idi/types.go",
		content: typesFile,
	},
	"validator.go": {
		path:    "internal/idi/validator.go",
		content: validatorFile,
	},
	"auth.go": {
		path:    "internal/idi/auth.go",
		content: authFile,
	},
	"rest.go": {
		path:    "internal/idi/rest.go",
		content: restFile,
	},
	"metrics.go": {
		path:    "internal/idi/metrics.go",
		content: metricsFile,
	},
	// framework files end
}

var dbFiles = map[string]*itemplate{
	"db.go": {
		path:    "internal/infrastructure/db.go",
		content: dbFile,
	},
}

var appDefaultFolders = map[string]string{
	"API":    "internal/apps/{app_name}/internal/adapters/{router_name}",
	"domain": "internal/apps/{app_name}/internal/application/domain",
	"ports":  "internal/apps/{app_name}/internal/application/ports",
	"tests":  "internal/apps/{app_name}/internal/tests",
}

var appDBFolders = map[string]string{
	"DB":   "internal/apps/{app_name}/internal/adapters/{db_name}",
	"sqlc": "internal/apps/{app_name}/internal/sqlc",
}

var appDBFiles = map[string]*itemplate{
	"irepository.go": {
		path:    "internal/apps/{app_name}/internal/application/ports/{app_name}.irepository.go",
		content: irepoFile,
	},
	"repository.go": {
		path:    "internal/apps/{app_name}/internal/adapters/{db_name}/{db_name}.repository.go",
		content: repoFile,
	},
}

var appDefaultFiles = map[string]*itemplate{
	"app.dtos.go": {
		path:    "internal/dtos/{app_name}.dtos.go",
		content: dtoFile,
	},
	"module.go": {
		path:    "internal/apps/{app_name}/app.init.go",
		content: initFile,
	},
	"domain.go": {
		path:    "internal/apps/{app_name}/internal/application/domain/{app_name}.domain.go",
		content: domainFile,
	},
	"app.go": {
		path:    "internal/apps/{app_name}/internal/application/{app_name}.app.go",
		content: appFile,
	},
	"controller.go": {
		path:    "internal/apps/{app_name}/internal/adapters/{router_name}/{app_name}.controller.go",
		content: ctrlrFile,
	},
	"routes.go": {
		path:    "internal/apps/{app_name}/internal/adapters/{router_name}/{app_name}.routes.go",
		content: routeFile,
	},
}

// app files below
// "app.dtos.go": "internal/apps/dtos/{app_name}.dtos.go",
// "module.go": "internal/apps/{app_name}/app.init.go",
// "domain.go": "internal/apps/{app_name}/application/domain/{app_name}.domain.go",
// "irepository.go": "internal/apps/{app_name}/application/ports/{app_name}.irepository.go",
// "repository.go": "internal/apps/{app_name}/internal/adapters/{db_name}/{db_name}.repository.go"
// "app.go": "internal/apps/{app_name}/application/{app_name}.app.go",
// "controller.go": "internal/apps/{app_name}/internal/adapters/{router_name}/{app_name}.controller.go",
// "routes.go": "internal/apps/{app_name}/internal/adapters/{router_name}/{app_name}.routes.go",

// framework files below
// "main.go": "cmd/{project_name}/main.go",
// "api.go": "cmd/{project_name}/api/api.go",
// "helpers.go": "cmd/{project_name}/api/helpers.go",
// "app.go": "cmd/{project_name}/api/{project_name}.app.go",

// "errors.go": "internal/idi/errors.go",
// "json.go": "internal/idi/json.go",
// "server.go": "internal/idi/server.go",
// "logger.go": "internal/idi/slogger.go",
// "types.go": "internal/idi/types.go",

// "dbname.go": "internal/infrastructure/{db_name}.go"
