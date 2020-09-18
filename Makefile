.PHONY: dev wire

dev:
	go run ./main.go 

wire:
	wire gen ./core/injector/wire.go
