# Learning Go language basics

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