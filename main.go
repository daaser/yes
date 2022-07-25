package main

import (
	"bufio"
	"os"
	"runtime/debug"
)

const (
	BUFSIZE    = 64 * 1024   // 64 Kib
	TOTAL_DATA = 1024 * 1024 // 1024 Kib * 64 Kib = 64 Gib
)

func partition(buf []byte, mid int) ([]byte, []byte) {
	return buf[0:mid], buf[mid:]
}

func fillBuffer(buffer, output []byte) []byte {
	if len(output) > len(buffer)/2 {
		return output
	}

	bufferSize := len(output)
	copy(buffer[:bufferSize], output)

	for ; bufferSize < len(buffer)/2; bufferSize *= 2 {
		left, right := partition(buffer, bufferSize)
		copy(right[:bufferSize], left)
	}
	return buffer[:bufferSize]
}

func main() {
	debug.SetGCPercent(-1)
	var (
		buffer [BUFSIZE]byte
		output = []byte{'y', '\n'}
		writer = bufio.NewWriter(os.Stdout)
	)

	filled := fillBuffer(buffer[:], output)
	for i := 0; i < TOTAL_DATA; i++ {
		writer.Write(filled)
	}
}
