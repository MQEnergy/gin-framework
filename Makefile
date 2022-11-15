APP=gin-framework

.PHONY: build
build:
	@go build -o releases/${APP}

.PHONY: windows
windows:
	@GOARCH=amd64 GOOS=windows go build -o releases/${APP}-windows

.PHONY: linux
linux:
	@GOARCH=amd64 GOOS=linux go build -o releases/${APP}-linux

.PHONY: darwin
darwin:
	@GOARCH=amd64 GOOS=darwin go build -o releases/${APP}-darwin

.PHONY: lint
lint:
	@golint ./...

.PHONY: generate
generate:
	@go generate -x

.PHONY: docker
docker:
	@docker build -t mqenergy/${APP}:latest .

.PHONY: clean
clean:
	@go clean -i .
	@rm -rf releases

.PHONY: help
help:
	@echo "2. make build - [go build]"
	@echo "3. make windows - [make window package]"
	@echo "4. make linux - [make linux package]"
	@echo "5. make darwin - [make darwin package]"
	@echo "6. make lint - [golint ./...]"
	@echo "7. make generate - [go generate -x]"
	@echo "8. make docker - [make docker images]"
	@echo "9. make clean - [remove releases files and cached files]"