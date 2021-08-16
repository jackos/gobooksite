.DEFAULT_GOAL := build
fmt:
	go fmt main.go
.PHONY:fmt
lint: fmt
	golint main.go
.PHONY:lint
vet: fmt
	go vet main.go
	shadow main.go
.PHONY:vet
build: vet
	go build main.go
.PHONY:build