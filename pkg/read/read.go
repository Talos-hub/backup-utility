package read

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type Data interface {
	~string | []byte
}

type statter interface {
	Stat() (os.FileInfo, error)
}

const (
	smallFileThreshold = 2 * 1024 * 1024
	bufferSize         = 128 * 1024
)

type Strategy int

const (
	StrategyAuto = iota
	StrategyBuffered
	StrategyDirect
	StrategyUntilNewLine
)

// Read reads data from reader with specified strategy
func Read[D Data](r io.Reader, strategy Strategy) (D, error) {
	var result []byte
	var zero D
	var err error

	// check stratedy and read
	switch strategy {
	case StrategyAuto:
		result, err = readAuto(r)
	case StrategyBuffered:
		result, err = readBuffered(r)
	case StrategyDirect:
		result, err = readDirect(r)
	case StrategyUntilNewLine:
		result, err = readUntilNewLine(r)
	default:
		return zero, fmt.Errorf("unknown strategy: %v", strategy)
	}

	if err != nil {
		return zero, err
	}

	return convertResult[D](result)
}

func convertResult[D Data](result []byte) (D, error) {
	var zero D

	switch any(zero).(type) {
	case string:
		return any(string(result)).(D), nil
	case []byte:
		return any(result).(D), nil
	default:
		return zero, fmt.Errorf("unsupported data type")
	}
}

// just simple read however it uses buffer
func readBuffered(r io.Reader) ([]byte, error) {
	reader := bufio.NewReaderSize(r, bufferSize)
	var result bytes.Buffer
	buffer := make([]byte, bufferSize)

	for {
		n, err := reader.Read(buffer)

		if n > 0 {
			result.Write(buffer[:n])
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("buffered reading error: %w", err)
		}

	}

	return result.Bytes(), nil

}

// A Simple funciton that read from Reader until newline
func readUntilNewLine(r io.Reader) ([]byte, error) {
	reader := bufio.NewReader(r)
	line, err := reader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("error reading: %w", err)
	}
	return line, nil
}

// If Reader is a file and it is large, then readbuffered is used,
// if not, then io.ReadAll is used.
func readAuto(r io.Reader) ([]byte, error) {
	// if r is file and size of the file is big
	// it use readbuffered()
	if file, ok := r.(statter); ok {
		info, err := file.Stat()
		if err != nil {
			return nil, fmt.Errorf("something is wrong, error getting file info: %w", err)
		}

		if info.Size() > smallFileThreshold {
			data, err := readBuffered(r)
			if err != nil {
				return nil, err
			}
			return data, nil
		}
	}

	data, err := readDirect(r)
	if err != nil {
		return nil, fmt.Errorf("error reading: %w", err)
	}
	return data, nil
}

// readDirect reads using io.ReadAll (direct read)
func readDirect(r io.Reader) ([]byte, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("direct reading error: %w", err)
	}
	return data, nil
}
