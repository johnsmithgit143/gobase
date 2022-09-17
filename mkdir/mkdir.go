package main

import (
	"github.com/johnsmithgit143/gobase/include"
	"flag"
	"fmt"
	"os"
)

const program string = "mkdir"

func usage() {
	fmt.Printf("usage: %s [-p] [-m mode] directory...\n\n", program)
	fmt.Printf("  -p        Creates any necessary parent directories.\n            No error if the target directory already exists.\n\n")
	fmt.Printf("  -m        Sets the permissions to be used when creating the directory.\n            The default is 0777.\n\n")
	fmt.Printf("  -h        Print help and exit.\n")
}

func main() {
	var (
		i int
		argv []string
		argc int
		fparent bool
		fmode int
		mode os.FileMode
		err error
	)
	
	flag.BoolVar(&fparent, "p", false, "")
	flag.IntVar(&fmode, "m", 0777, "")
	flag.Usage = usage
	flag.Parse()
	
	argv = flag.Args()
	argc = len(argv)
	
	if argc == 0 || fmode > 0777 {
		flag.Usage()
		os.Exit(1)
	}
	
	mode = os.FileMode(fmode)
	
	for i = 0; i < argc; i++ {
		if fparent {
			if err = os.MkdirAll(argv[i], mode); err != nil {
				include.Sysfatal(program, err)
			}
		} else {
			if err = os.Mkdir(argv[i], mode); err != nil {
				include.Sysfatal(program, err)
			}
		}
	}
}

