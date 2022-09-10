package main

import (
	"github.com/johnsmithgit143/gobase/include"
	"github.com/johnsmithgit143/gobase/flag"
	"fmt"
)

func main() {
	validflags = []string{
		"-p", "-m",
	}
	fmt.Println(flag.Parse(validFlags))
}
