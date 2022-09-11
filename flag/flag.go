package flag

import (
	"os"
	"fmt"
	"strings"
)

func Parse(validFlags []string) (map[string]int, error) {
	var (
		flags = make(map[string]int)
		i int
		x int
	)
	START:
	for i = 1; i < len(os.Args); i++ {
		if os.Args[i] == "--" {
			break
		}
		if !strings.HasPrefix(os.Args[i], "-") {
			continue
		}
		for x = 0; x < len(validFlags); x++ {
			if os.Args[i] == validFlags[x] {
				flags[os.Args[i]] = i
				continue START
			}
		}
		return flags, fmt.Errorf("invalid option: %s", os.Args[i])
	}
	return flags, nil
}
