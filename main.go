package main

import (
	"bufio"
	"os"
	"runtime/debug"
)

const (
	BUFSIZE    = 64 * 1024       // 64 Kib
	TOTAL_DATA = 2 * 1024 * 1024 // 64 Gib
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
		output = parseArgs()
		writer = bufio.NewWriterSize(os.Stdout, BUFSIZE)
	)

	filled := fillBuffer(buffer[:], output)
	for i := 0; i < TOTAL_DATA; i++ {
		writer.Write(filled)
	}
}

func parseArgs() []byte {
	if len(os.Args) == 2 {
		return append([]byte(os.Args[1]), '\n')
	} else {
		return []byte{'y', '\n'}
	}
}
