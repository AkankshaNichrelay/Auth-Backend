APP = auth-backend
PKG = ...

deps:
	go mod tidy -compat=1.17

wire:
	@go get -d github.com/google/wire/cmd/wire
	@cd cmd/$(APP) && wire

test:
	@go test -p 1 ./$(PKG)

build:
	@echo "building app"
	@go build .\cmd\$(APP)

run:
	@.\$(APP).exe
