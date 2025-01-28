package todos_test

//go:generate mockgen -source ..\application\ports\todos.irepository.go -destination .\repo.mock.go -package todos_test
//go:generate mockgen -source .\generate_test.go -destination .\app.mock.go -package todos_test
