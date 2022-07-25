package main

import (
	"bufio"
	"os"
)

const BUFSIZE = 64 * 1024 // 64 Kib

func main() {
	writer := bufio.NewWriter(os.Stdout)
	for {
		writer.Write([]byte("y\n"))
	}
}
