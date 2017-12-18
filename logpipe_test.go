package logpipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testMatch(string) bool {
	return true
}

func testDontMatch(string) bool {
	return false
}

func TestNew(t *testing.T) {
	_, err := New("some/file/path")
	assert.Nil(t, err, "New log pipe can be created")
}

func TestPath(t *testing.T) {
	p, _ := New("some/file/path")
	assert.Equal(t, "some/file/path", p.Path(), "Path should be set as passed")
}

func TestLogPipeDefaultMatch(t *testing.T) {
	p, _ := New("some/file/path")
	assert.True(t, p.Match("arbitrary line"), "Match should be true")
}

func TestLogPipeMatch(t *testing.T) {
	p, _ := New("some/file/path")
	p.MatchFunc(testMatch)
	assert.True(t, p.Match("arbitrary line"), "Match should be true")
}

func TestLogPipeNoMatch(t *testing.T) {
	p, _ := New("some/file/path")
	p.MatchFunc(testDontMatch)
	assert.False(t, p.Match("arbitrary line"), "Match should be false")
}

func TestLogPipeHandle(t *testing.T) {
	var recorder string
	handle := func(l string) error {
		recorder = l
		return nil
	}
	expected := "arbitrary line"

	p, _ := New("some/file/path")
	p.MatchFunc(testMatch)
	p.HandleFunc(handle)

	assert.True(t, p.Match(expected), "Match should be true")
	err := p.Handle(expected)

	assert.Nil(t, err, "Line should be handled")
	assert.Equal(t, expected, recorder, "Recorder should match expected line")
}

func TestLogPipeHandleNoHandler(t *testing.T) {
	appliedMatch := false
	matcher := func(line string) bool {
		appliedMatch = true
		return true
	}

	p, _ := New("some/file/path")
	p.MatchFunc(matcher)

	err := p.Handle("something")

	assert.Nil(t, err, "Handler should not return an error")
	assert.False(t, appliedMatch, "Match function should not be tried if no handler exists")
}

func TestLogPipeHandleNoMatch(t *testing.T) {
	var recorder string
	handle := func(l string) error {
		recorder = l
		return nil
	}
	expected := "arbitrary line"

	p, _ := New("some/file/path")
	p.MatchFunc(testDontMatch)
	p.HandleFunc(handle)

	assert.False(t, p.Match(expected), "Match should be false")
	err := p.Handle(expected)

	assert.Nil(t, err, "Handler should not return error")
	assert.Empty(t, recorder, "Recorder should be empty")
}
