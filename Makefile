## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## build_cli:builds the command line tool goravel and copies it to myapp
build_cli:
	@go build -o ../myapp/goravel ./cmd/cli