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