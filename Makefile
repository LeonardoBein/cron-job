.PHONY: all
all: build
FORCE: ;

SHELL  := env LANGUAGE_ENV=$(LANGUAGE_ENV) $(SHELL)
LANGUAGE_ENV ?= ptBR

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-app

build-app: 
	go build -tags $(LANGUAGE_ENV) -o ./bin/cron-service main.go

linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(LANGUAGE_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/cron-service main.go
