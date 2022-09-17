package main

import (
	"github.com/johnsmithgit143/gobase/include"
	"bufio"
	"os"
)

const program string = "read"

func main() {
	var (
		scanner *bufio.Scanner
		err error
	)
	
	scanner = bufio.NewScanner(os.Stdin)
	
	scanner.Scan()
	
	if err = scanner.Err(); err != nil {
		include.Sysfatal(program, err)
	}
	
	if _, err = os.Stdout.Write(scanner.Bytes()); err != nil {
		include.Sysfatal(program, err)
	}
}