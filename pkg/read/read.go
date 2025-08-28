package read

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Data interface {
	~string | []byte
}

const (
	smallFileThreshold = 2 * 1024 * 1024
	buffered           = 128 * 1024
)

type Strategy int

const (
	StrategyAuto = iota
	StrategyBuffered
	StrategyDirect
)

// Read read a data from reader
// I has a few strategy of reading
// StrategyAuto, StrateguBufferd and StrategyDirect
func Read[D Data](r io.Reader, strategy Strategy) (D, error) {
	var result []byte
	var data D
	var err error

	// check stratedy and read
	switch strategy {
	case StrategyAuto:
		result, err = readAuto(r)
	case StrategyBuffered:
		result, err = readBufferd(r)
	case StrategyDirect:
		result, err = readBufferd(r)
	}

	if err != nil {
		return data, nil
	}

	switch any(data).(type) {
	case string:
		return D(string(result)), nil
	case []byte:
		return D(result), nil
	default:
		return data, nil
	}
}

// just simple read however it uses buffer
func readBufferd(r io.Reader) ([]byte, error) {
	reader := bufio.NewReaderSize(r, buffered)
	var result []byte

	for {
		chunk := make([]byte, buffered)

		n, err := reader.Read(chunk)

		if n > 0 {
			result = append(result, chunk...)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("something is wrong, error buffered reading: %w", err)
		}
	}

	return result, nil

}

// If Reader is a file and it is large, then readbuffered is used,
// if not, then io.ReadAll is used.
func readAuto(r io.Reader) ([]byte, error) {
	// if r is file and size of the file is big
	// it use readbuffered()
	if file, ok := r.(*os.File); ok {
		info, err := file.Stat()
		if err != nil {
			return nil, fmt.Errorf("something is wrong, error getting file info: %w", err)
		}

		if info.Size() > smallFileThreshold {
			data, err := readBufferd(r)
			if err != nil {
				return nil, err
			}
			return data, nil
		}
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("error reading: %w", err)
	}
	return data, nil
}
