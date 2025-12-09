# What is this?
This is a simple Brainf_ck interpreter written in Go. I wrote it in an evening because I was bored. A "Hello world" program file is also included in the `cmd/` folder.

# Building
```bash
go build -o foo.exe
```

# Usage
## REPL
To launch REPL mode, either run the compiled application in the terminal, or run `cmd/main.go`.

## Reading a File
To execute a program from a file, you need to pass it as an argument when starting the application.
Example:
```bash
go run ./cmd/main.go foo.bf
```
