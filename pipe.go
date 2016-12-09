package pipe

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// used here to prevent os package import
var (
	Stdin  = os.Stdin
	Stdout = os.Stdout
	Stderr = os.Stderr
)

// StdDataChannel used for as async communication channel
type StdDataChannel chan StdData

// StdData is used as channel structure
type StdData struct {
	Data []byte
	Err  error
}

// isNamedPipe checks whether valid pipe used
func isNamedPipe() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return ((fileInfo.Mode() & os.ModeNamedPipe) == os.ModeNamedPipe)
}

// readChunk read from io.reader amount of bytes defined in size
// returns error if something is wrong or []byte if its fine
func readChunk(r io.Reader, size int) ([]byte, error) {
	buf := make([]byte, 0, size)
	n, err := r.Read(buf[:cap(buf)])
	if err != nil {
		return []byte{}, err
	}
	// fallback for EOF
	if n == 0 {
		return []byte{}, io.EOF
	}
	return buf[:n], nil
}

// Read reads from io.reader one chunk as defined as bufSize
func Read(pipe io.Reader, bufSize int) ([]byte, error) {
	if !isNamedPipe() {
		return []byte{}, errors.New("Not Valid Named Pipe")
	}
	return readChunk(bufio.NewReader(pipe), bufSize)
}

// AsyncRead keeps reading from io.reader into buffer
// return data in stdData structure
// use in only with routine - ie go AsyncRead
func AsyncRead(pipe io.Reader, bufSize int, stdDataChan chan StdData) {
	if !isNamedPipe() {
		stdDataChan <- StdData{
			Err: errors.New("Not Valid Named Pipe"),
		}
	}
	r := bufio.NewReader(pipe)
	for {
		stdData := StdData{}
		data, err := readChunk(r, bufSize)
		if err != nil {
			stdData.Err = err
		}
		stdData.Data = data
		stdDataChan <- stdData
	}
}
