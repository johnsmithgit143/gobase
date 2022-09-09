package include

import (
	"os"
	"fmt"
)

func Sysfatal(program string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", program, err.Error())
	os.Exit(1)
}
