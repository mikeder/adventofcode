.DEFAULT_GOAL := help

test:
	go test -coverprofile=cov.out ./...
	go tool cover -html=cov.out -o cov.html
doc:
	go doc ./...
run:
	go run main.go
help:
	$(MAKEFILE_LIST)
