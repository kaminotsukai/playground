package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	filePath := "example.mp4"
	e := hasInvalidNALUnitSize(filePath)
	fmt.Println(e)
}

func hasInvalidNALUnitSize(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data)-4; i++ {
			if data[i] == 0 && data[i+1] == 0 && data[i+2] == 0 && data[i+3] == 1 {
				if i > 0 && data[i-1] == 0 {
					start := i + 4
					if start >= len(data) {
						return 0, nil, nil
					}
					naluSize := binary.BigEndian.Uint32(data[start-4 : start])
					end := start + int(naluSize)
					if end >= len(data) {
						return len(data), data, nil
					}
					return end, data[start:end], nil
				}
			}
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	for scanner.Scan() {
		if isInvalidNALUnitSize(scanner.Bytes()) {
			return true
		}
	}
	return false
}

func isInvalidNALUnitSize(data []byte) bool {
	for i := 0; i < len(data)-4; i++ {
		if data[i] == 0 && data[i+1] == 0 && data[i+2] == 0 && data[i+3] == 1 {
			if i > 0 && data[i-1] == 0 {
				start := i + 4
				if start >= len(data) {
					return false
				}
				naluSize := binary.BigEndian.Uint32(data[start-4 : start])
				end := start + int(naluSize)
				if end >= len(data) {
					return true
				}
			}
		}
	}
	return false
}
