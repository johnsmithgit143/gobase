package main

import (
	"github.com/johnsmithgit143/gobase/flag"
	"github.com/johnsmithgit143/gobase/include"
	"fmt"
	"os"
)

const program string = "mkdir"

func main() {
	var (
		validFlags = []string{
			"-p", "-m",
		}
		flags map[string]int
		parent bool
		i int
		ok bool
		err error
	)
	if flags, err = flag.Parse(validFlags); err != nil {
		include.Sysfatal(program, err)
	}
	if _, ok = flags["-p"]; ok {
		parent = true
	}
	fmt.Println(parent)

	for i = 1; i < len(os.Args); i++ {
		if parent {
			if err = os.MkdirAll(os.Args[1], 0777); err != nil {
				include.Sysfatal(program, err)
			}
		}
	}

}
