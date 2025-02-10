.PHONY: generate createsql testv testdox migrateup migratedown

run:
	go run .\...

generate:
	go generate .\...
	go mod tidy
	go fmt .\...

testv:
	go test .\... -count=1 -race --shuffle on -v 

testdox:
	gotestdox .\... -count=1 -race --shuffle on -v 

createsql:
	goose -s -dir .\migrations create ${t} sql

