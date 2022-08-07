generate:
	wire ./src/injector

run:
	go run main.go

test:
	go test ./src/services -v

swagger:
	swag init

all: generate swagger run