package main

import (
	"os"
	"fmt"
	"strings"
)

var Validflags = [...]string{
	"-a", "-b", "-c", "--testing",
}

func Parse() (map[string]int, error) {
	var (
		flags = make(map[string]int)
		i int
		x int
	)
	start:
	for i = 1; i < len(os.Args); i++ {
		if os.Args[i] == "--" {
			break
		}
		if !strings.HasPrefix(os.Args[i], "-") {
			continue
		}
		for x = 0; x < len(Validflags); x++ {
			if os.Args[i] == Validflags[x] {
				flags[os.Args[i]] = i
				continue start
			}
		}
		return flags, fmt.Errorf("invalid option: %s", os.Args[i])
	}
	return flags, nil
}

func foo() {
	START:
	for y := 0; y < 5; y++ {
		fmt.Println("y:", y)
		for x := 0; x < 5; x++ {
			if x == 4 {
				continue START
			}
			fmt.Println("x:", x)
		}
	}
}

func main() {
	foo()
	fmt.Println(Parse())
}