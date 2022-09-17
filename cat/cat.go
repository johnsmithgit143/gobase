package main

import (
	"fmt"
	"os"
	"io"
	"github.com/johnsmithgit143/gobase/include"
)

const program string = "cat"

func main() {
	var (
		f *os.File
		argc int = len(os.Args)
		i int
		err error
	)
	
	if argc == 1 {
		fmt.Fprintln(os.Stderr, "usage: cat [files...]")
		os.Exit(1)
	}
	
	for i = 1; i < argc; i++ {
		if f, err = os.Open(os.Args[i]); err != nil {
			include.Sysfatal(program, err)
		}
		
		if _, err = io.Copy(os.Stdout, f); err != nil {
			include.Sysfatal(program, err)
		}
		
		if err = f.Close(); err != nil {
			include.Sysfatal(program, err)
		}
	}
}
