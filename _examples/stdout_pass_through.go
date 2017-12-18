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
