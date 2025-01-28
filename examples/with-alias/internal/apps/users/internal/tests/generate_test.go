package todos_test

//go:generate mockgen -source ..\application\ports\users.irepository.go -destination .\repo.mock.go -package users_test
//go:generate mockgen -source .\generate_test.go -destination .\app.mock.go -package users_test
