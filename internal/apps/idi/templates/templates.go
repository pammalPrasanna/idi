package templates

type itemplate struct {
	path    string
	content []byte
}

var frameworkFolders = map[string]string{
	"cmd":   "cmd/api",
	"apps":  "internal/apps",
	"dtos":  "internal/dtos",
	"lib":   "internal/lib",
	"tests": "internal/tests",

	"infrastructure": "internal/infrastructure",
}

var frameworkDefaultFiles = map[string]*itemplate{
	// cmd files start
	"api.go": {
		path:    "cmd/api/api.go",
		content: apiFile,
	},
	"main.go": {
		path:    "main.go",
		content: mainFile,
	},
	// cmd files ends

	// framework files start
	"configs.go": {
		path:    "internal/lib/configs.go",
		content: configsFile,
	},
	"errors.go": {
		path:    "internal/lib/errors.go",
		content: errorsFile,
	},
	"json.go": {
		path:    "internal/lib/json.go",
		content: jsonFile,
	},
	"server.go": {
		path:    "internal/lib/server.go",
		content: serverFile,
	},
	"logger.go": {
		path:    "internal/lib/slogger.go",
		content: loggerFile,
	},
	"types.go": {
		path:    "internal/lib/types.go",
		content: typesFile,
	},
	"validator.go": {
		path:    "internal/lib/validator.go",
		content: validatorFile,
	},

	"rest.go": {
		path:    "internal/lib/rest.go",
		content: restFile,
	},
	"metrics.go": {
		path:    "internal/lib/metrics.go",
		content: metricsFile,
	},
	"idi.go": {
		path:    "internal/lib/{alias}.go",
		content: idiFile,
	},
	"middlewares.go": {
		path:    "internal/lib/middlewares.go",
		content: middlewaresFile,
	},
	"itime.go": {
		path:    "internal/lib/itime.go",
		content: itimeFile,
	},
	".env": {
		path:    ".env",
		content: dotEnvFile,
	},
	"makeFile": {
		path:    "Makefile",
		content: makeFile,
	},
	// framework files end
}

var dbFiles = map[string]*itemplate{
	"db.go": {
		path:    "internal/infrastructure/db.go",
		content: dbFile,
	},
}

var authFolders = map[string]string{
	"auth": "internal/lib/auth",
}

var jwtFiles = map[string]*itemplate{
	"imaker.go": {
		path:    "/internal/lib/auth/imaker.go",
		content: imakerFile,
	},
	"auth.go": {
		path:    "internal/lib/auth.go",
		content: authFile,
	},
	"jwt_maker.go": {
		path:    "/internal/lib/auth/jwt_maker.go",
		content: jwtMakerFile,
	},
}

var pasetoFiles = map[string]*itemplate{
	"imaker.go": {
		path:    "/internal/lib/auth/imaker.go",
		content: imakerFile,
	},
	"auth.go": {
		path:    "internal/lib/auth.go",
		content: authFile,
	},
	"paseto_maker.go": {
		path:    "/internal/lib/auth/paseto_maker.go",
		content: pasetoMakerFile,
	},
}

var appDefaultFolders = map[string]string{
	"API":    "internal/apps/{app_name}/internal/adapters/{router_name}",
	"domain": "internal/apps/{app_name}/internal/application/domain",
	"ports":  "internal/apps/{app_name}/internal/application/ports",
	"tests":  "internal/apps/{app_name}/internal/tests",
}

var appDBFolders = map[string]string{
	"DB": "internal/apps/{app_name}/internal/adapters/{db_name}",
	"migrations": "migrations",
	// "sqlc": "internal/apps/{app_name}/internal/sqlc",
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
	"users.sql": {
		path: "migrations/00001_create_users_table.sql",
		content: usersMigrationFile,
	},
}

var appDefaultFiles = map[string]*itemplate{
	"app.dtos.go": {
		path:    "internal/dtos/{app_name}.dtos.go",
		content: dtoFile,
	},
	"http.dtos.go": {
		path:    "internal/dtos/http.dtos.go",
		content: httpDTOsFile,
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
	"rest.dtos.go": {
		path:    "internal/apps/{app_name}/internal/adapters/{router_name}/rest.dtos.go",
		content: restDTOsFile,
	},

	// 
	"generate_test.go": {
		path:    "internal/apps/{app_name}/internal/tests/generate_test.go",
		content: genTestFile,
	},
	"repository_test.go": {
		path:    "internal/apps/{app_name}/internal/tests/repository_test.go",
		content: repoTestFile,
	},
	"main_test.go": {
		path:    "internal/apps/{app_name}/internal/tests/main_test.go",
		content: mainTestFile,
	},
	"application_test.go": {
		path:    "internal/apps/{app_name}/internal/tests/application_test.go",
		content: applicationTestFile,
	},
	"controller_test.go": {
		path:    "internal/apps/{app_name}/internal/tests/controller_test.go",
		content: controllerTestFile,
	},
	"random_test.go": {
		path:    "internal/apps/{app_name}/internal/tests/random_test.go",
		content: randomTestFile,
	},
	"schema_test.go": {
		path:    "internal/apps/{app_name}/internal/tests/schema_test.go",
		content: schemaTestFile,
	},
	"testdb_test.go": {
		path:    "internal/apps/{app_name}/internal/tests/testdb_test.go",
		content: testDBTestFile,
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
// "api.go": "cmd/api/api.go",
// "helpers.go": "cmd/api/helpers.go",
// "app.go": "cmd/api/{project_name}.app.go",

// "errors.go": "internal/lib/errors.go",
// "json.go": "internal/lib/json.go",
// "server.go": "internal/lib/server.go",
// "logger.go": "internal/lib/slogger.go",
// "types.go": "internal/lib/types.go",

// "dbname.go": "internal/infrastructure/{db_name}.go"
