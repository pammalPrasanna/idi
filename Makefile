.PHONY: generate-examples

generate-examples:
	mkdir examples
	go run .\... -cp with-alias -cdb sqlite3 -ca users -auth -paseto -a rootApp
	go run .\... -cp without-alias -cdb postgres -ca users -auth -paseto
