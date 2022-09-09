package main

import (
	"fmt"
	"os"
	"io"
)

func sysfatal(err error) {
	fmt.Fprintln(os.Stderr, "cat:", err)
	os.Exit(1)
}

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
	} else {
		for i = 1; i < argc; i++ {
			if f, err = os.Open(os.Args[i]); err != nil {
				sysfatal(err)
			}
			if _, err = io.Copy(os.Stdout, f); err != nil {
				sysfatal(err)
			}
			if err = f.Close(); err != nil {
				sysfatal(err)
			}
		}
	}
}
