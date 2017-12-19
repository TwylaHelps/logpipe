package logpipe

import (
	"bufio"
	"os"
	"syscall"
)

func removeExisting(path string) error {
	return os.Remove(path)
}

func mkFifo(path string) error {
	return syscall.Mkfifo(path, 0666)
}

func getFileReader(path string) (*bufio.Reader, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(file), nil
}
