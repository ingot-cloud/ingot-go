.PHONY: dev wire

VERSION = 0.1.0
BIN_PATH = ./cmd/bin/ingot

build:
	go build -ldflags "-w -s -X main.VERSION=$(VERSION)" -o $(BIN_PATH) ./cmd/ingot

dev:
	go run -ldflags "-X main.VERSION=$(VERSION)" ./cmd/ingot/main.go s -c ./configs/config.yml -m ./configs/casbin_model.conf

watch:
	air -c .air.conf

wire:
	wire gen ./internal/app/core/injector/wire.go
