generate:
	wire ./src/injector

run:
	go run main.go

all: generate run