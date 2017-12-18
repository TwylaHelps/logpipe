// Package logpipe is a middleware between processes that log to files and
// metrics as well as log aggregation services that expect a certain format.
package logpipe

import (
	"fmt"
)

// LogPipe is the API interface to manage and run the matching and handling
// logic.
type LogPipe interface {
	Match(string) bool
	Handle(string) error
	MatchFunc(func(string) bool)
	HandleFunc(func(string) error)
	Path() string
	Run() error
}

type logPipe struct {
	matchFunc  func(string) bool
	handleFunc func(string) error
	path       string
}

func defaultMatchFunc(string) bool {
	return true
}

// Match applies the function passed in with MatchFunc(). If none was added
// before the first match, the default match function will be used that just
// returns true and thus matches every input.
func (l *logPipe) Match(line string) bool {
	return l.matchFunc(line)
}

// Handle applies the function passed in with HandleFunc() in case Match()
// returns true. If no function to handle a line was set no match will be
// tried and Handle returns nil.
func (l *logPipe) Handle(line string) error {
	if l.handleFunc == nil {
		return nil
	}
	if l.Match(line) {
		return l.handleFunc(line)
	}
	return nil
}

// MatchFunc sets the function used to determine if the handler function
// should be applied to the line.
func (l *logPipe) MatchFunc(mf func(string) bool) {
	l.matchFunc = mf
}

// HandleFunc sets the function that will be applied to a line if the matcher
// function for that line is true.
func (l *logPipe) HandleFunc(hf func(string) error) {
	l.handleFunc = hf
}

// Path returns the path of the pipe that is read from be this LogPipe.
func (l *logPipe) Path() string {
	return l.path
}

// Run creates the actual named pipe and starts consuming lines to match and
// handle. In case basic setup of the pipe fails Run will return an error,
// otherwise it will block waiting for incoming lines.
func (l *logPipe) Run() error {
	// Remove the pipe or file if any exists
	err := removeExisting(l.path)
	if err != nil {
		return err
	}
	err = mkFifo(l.path)
	if err != nil {
		return err
	}
	reader, err := getFileReader(l.path)
	if err != nil {
		return err
	}

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			continue
		}
		err = l.Handle(string(line))
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

// New returns a new instance that implements LogPipe. The desired path to the
// named pipe and the relevant fields have to be passed in.
//
// NOTE: the path will be removed if it exists!
func New(path string) (LogPipe, error) {
	lp := &logPipe{
		path: path,

		matchFunc: defaultMatchFunc,
	}
	return lp, nil
}
