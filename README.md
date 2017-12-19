# LogPipe

LogPipe creates a named pipe and listens line wise to all input. The use case in
mind is a middleware between processes that log to files and metrics as well as
log aggregation services.

## Install

    go get github.com/TwylaHelps/logpipe

## Usage

Create and listen to a pipe:

    package main

    import (
        "fmt"

        "github.com/twylahelps/logpipe"
    )

    func main() {
        pipe, err := logpipe.New("/tmp/my-nice-pipe")
        if err != nil {
            panic(err)
        }
        counter := 0
        handler := func(line string) error {
            counter++
            fmt.Printf("%d %s", counter, line)
            return nil
        }
        pipe.HandleFunc(handler)
        err = pipe.Run()
        if err != nil {
            fmt.Println("Error:", err)
        }
    }

Run:

    go run _examples/stdout_pass_through.go

In a different terminal echo something to the named pipe:

    echo this is a test >> /tmp/my-nice-pipe
