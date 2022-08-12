APP=gin-framework

bindata:
	@go get -u github.com/go-bindata/go-bindata/...
	@go install github.com/go-bindata/go-bindata/...
	@go-bindata -o=./config/config_yaml.go -pkg=config config.*.yaml
build:
	@go-bindata -o=./config/config_yaml.go -pkg=config config.*.yaml
	@go build -o releases/${APP}
windows:
	@go-bindata -o=./config/config_yaml.go -pkg=config config.*.yaml
	@GOARCH=amd64 GOOS=windows go build -o releases/${APP}-windows
linux:
	@go-bindata -o=./config/config_yaml.go -pkg=config config.*.yaml
	@GOARCH=amd64 GOOS=linux go build -o releases/${APP}-linux
darwin:
	@go-bindata -o=./config/config_yaml.go -pkg=config config.*.yaml
	@GOARCH=amd64 GOOS=darwin go build -o releases/${APP}-darwin
lint:
	@golint ./...
generate:
	@go generate -x
docker:
	@docker build -t lyky-tech/${APP}:latest .
clean:
	@go clean -i .
	@rm -rf releases
help:
	@echo "1. make bindata - [go-bindata]"
	@echo "2. make build - [go-bindata + go build]"
	@echo "3. make windows - [go-bindata + make window package]"
	@echo "4. make linux - [go-bindata + make linux package]"
	@echo "5. make darwin - [go-bindata + make darwin package]"
	@echo "6. make lint - [golint ./...]"
	@echo "6. make generate - [go generate -x]"
	@echo "7. make docker - [make docker images]"
	@echo "8. make clean - [remove releases files and cached files]"