package logpipe

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTempFileName() string {
	// Create and remove a temp file to make rather sure a file with that
	// name does not yet exist.
	temp, _ := ioutil.TempFile("", "")
	name := temp.Name()
	temp.Close()
	os.Remove(name)

	return name
}

func TestRemoveExisting(t *testing.T) {
	temp, _ := ioutil.TempFile("", "")
	name := temp.Name()
	// Remove in case the tests failed early
	defer os.Remove(name)

	// Just checking if the setup worked so far
	_, err := os.Stat(name)
	assert.Nil(t, err, "Temporary file should exist")

	err = removeExisting(name)
	assert.Nil(t, err, "Removing the file should not cause errors")

	_, err = os.Stat(name)
	assert.True(t, os.IsNotExist(err), "Temporary file should be removed")
}

func TestMkFifo(t *testing.T) {
	name := getTempFileName()
	_, err := os.Stat(name)
	assert.True(t, os.IsNotExist(err), "File with desired name should should not exist")
	err = mkFifo(name)
	defer os.Remove(name)
	assert.Nil(t, err, "Creating the named pipe should not cause errors")

	_, err = os.Stat(name)
	assert.Nil(t, err, "The pipe should exist")
}

func TestGetFileReader(t *testing.T) {
	name := getTempFileName()
	_, err := os.Stat(name)
	assert.True(t, os.IsNotExist(err), "File with desired name should should not exist")
	err = mkFifo(name)
	defer os.Remove(name)
	assert.Nil(t, err, "Creating the named pipe should not cause errors")

	// TODO: hangs, waiting for feedback from #go-nuts
	//_, err = getFileReader(name)
	//assert.Nil(t, err, "Getting the pipe reader should not cause errors")
}
