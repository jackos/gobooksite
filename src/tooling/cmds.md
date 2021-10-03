# Cmds
## Upgrade version of module
go get -u=patch github.com/[user]/[repo]

## List variables
go env

## See where the overwrite file is for GO
go env GOENV


## Imports sorter
```fish
go install golang.org/x/tools/cmd/goimports@latest

goimports -l -w .
```
The -l flag tells goimports to print the files with incorrect formatting to the console. The -w flag tells goimports to modify the files in-place. The . specifies the files to be scanned: everything in the current directory and all of its subdirectories.

## Go Lint
```fish
go install golang.org/x/lint/golint@latest

# run it with:
golint ./...
```

## Go vet
Passing wrong numbers of paramteres, assigning values that are never used etc
```fish
go vet ./...
```

## golangci-lint
Many different tools run over the same directory
```fish
golangci-lint run
```

## shadow detection
```fish
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
```

## makefile to fmt, lint, vet and build
```makefile
.DEFAULT_GOAL := build

fmt:
	go fmt main.go
.PHONY:fmt
lint: fmt
	golint main.go
.PHONY:lint
vet: fmt
	go vet main.go
.PHONY:vet
build: vet
	go build main.go
.PHONY:build
```

## Use a different version of go
```fish
go get golang.org/dl/go1.15.6
go1.15.6 download
```
## Delete new version of go
```fish
go1.15.6 env GOROOT
rm -rf $(go1.15.6 env GOROOT)
rm $(go env GOPATH)/bin/go1.15.6
```