package main

import (
	"bufio"
	"os"
)

const BUFSIZE = 64 * 1024 // 64 Kib

func partition(buf []byte, mid int) ([]byte, []byte) {
	return buf[0:mid], buf[mid:]
}

func fillBuffer(buffer, output []byte) []byte {
	if len(output) > len(buffer)/2 {
		return output
	}

	bufferSize := len(output)
	copy(buffer[:bufferSize], output)

	for bufferSize < len(buffer)/2 {
		left, right := partition(buffer, bufferSize)
		copy(right[:bufferSize], left)
		bufferSize *= 2
	}
	return buffer[:bufferSize]
}

func main() {
	var (
		buffer [BUFSIZE]byte
		output = []byte("y\n")
	)

	filled := fillBuffer(buffer[:], output)
	writer := bufio.NewWriter(os.Stdout)
	for {
		writer.Write(filled)
	}
}
