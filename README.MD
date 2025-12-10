# What is this?
This is a simple Brainf_ck interpreter written in Go. I wrote it in an evening because I was bored. A "Hello world" program file is also included in the `cmd/` folder.

# Building
```bash
cd ./cmd
go build -o bf
#or
go build -o bf.exe
```
After you can add path to the `./cmd` folder into the $PATH to use it in anyware

You also can change name `bf` to any other, it's up to you

# Usage
## REPL
To launch REPL mode, either run the compiled application in the terminal, or run `cmd/main.go`.

## Reading a File
To execute a program from a file, you need to pass it as an argument when starting the application.
Example:
```bash
go run ./cmd/main.go foo.bf

#if you build u can do this instead
bf foo.bf
```

To start REPL you can do this:
```bash
go run ./cmd/main.go
# or just
bf
```
