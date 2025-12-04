package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
)

func SkipPrefix(body *bufio.Reader) error {
	var isPrefix bool
	var err error
	for i := 0; i < 2; i++ {
		_, isPrefix, err = body.ReadLine()
		if err != nil || isPrefix {
			return fmt.Errorf("error when reading response with the serialized city names: %w", err)
		}
	}
	return nil
}

func ReadLine(body *bufio.Reader) ([]byte, error) {
	bytesToDecode, isPrefix, err := body.ReadLine()
	if err != nil {
		return nil, err
	}
	if !isPrefix {
		return bytesToDecode, nil
	}
	bytesToDecodeFinal := make([]byte, len(bytesToDecode))
	copy(bytesToDecodeFinal, bytesToDecode)
	for isPrefix {
		bytesToDecode, isPrefix, err = body.ReadLine()
		if err != nil {
			return bytesToDecode, err
		}
		bytesToDecodeFinal = append(bytesToDecodeFinal, bytesToDecode...)
	}
	return bytesToDecodeFinal, nil
}

func GetInnerBytes(body *bufio.Reader) ([]byte, error) {
	bytesToDecode, err := ReadLine(body)
	if err != nil {
		return nil, err
	}

	innerBytes := ""

	err = json.Unmarshal(bytesToDecode, &[][]interface{}{{nil, nil, &innerBytes}})
	if err != nil {
		return nil, fmt.Errorf("error during decoding master schema: %v", err)
	}
	return []byte(innerBytes), nil
}
