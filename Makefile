## Build all binaries

build-ch01:
	go build -o bin/http-go-server ch01/cmd/server/main.go

compile-ch02-pb:
	protoc ch02/api/v1/*.proto --go_out=. --go_opt=paths=source_relative --proto_path=.
test:
	go test -race ./...