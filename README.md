# Learning Go language basics

## IDEs to use

* Currently, I'm working within IntelliJ IDEA with Go plugin installed.
  * I haven't yet installed the completely separate GoLand IDE. 
* It's also working in VS Code with Go extension (provided by the Go team itself!) and installed Go tools.

## Create a module

Module name is `demo`

```bash
go mod init demo
```

## Run the main class

Just this doesn't work, we need to specify what to run:
```bash
go run
```

These ones work:
```bash
go run .

go run ./main.go

go run main.go
`````

## Get a dependency

```bash
go get golang.org/x/exp/slices
```

## Run tests

Run test in a package `testing`:
```bash
go testt testing
```

Run all tests in current directory and all subdirectories (subpackages):

```bash
go test ./...
```

## Useful documentation links

* https://pkg.go.dev/builtin@go1.26.4 — built-in types
* https://pkg.go.dev/std — standard library
* https://go.dev/ref/spec — language specification
* https://go.dev/wiki/Iota — `iota` explanation
* https://go.dev/doc/effective_go — Effective Go
* https://go-proverbs.github.io/ — Go Proverbs (they link to the video where they were introduced)
* https://pkg.go.dev/golang.org/x/exp/constraints — a package with built-in type constraints
* https://pkg.go.dev/golang.org/x/exp/slices — a package with generic slices operations
* https://pkg.go.dev/golang.org/x/exp/maps — a package with generic maps operations
* https://pkg.go.dev/testing — testing package
* https://go.dev/doc/diagnostics — performance tests