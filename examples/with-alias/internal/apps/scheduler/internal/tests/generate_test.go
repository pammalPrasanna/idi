package todos_test

//go:generate mockgen -source ..\application\ports\scheduler.irepository.go -destination .\repo.mock.go -package scheduler_test
//go:generate mockgen -source .\generate_test.go -destination .\app.mock.go -package scheduler_test
