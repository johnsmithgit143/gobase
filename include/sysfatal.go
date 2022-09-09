package include

import (
	"os"
	"fmt"
)

func Sysfatal(program string, err error) {
	fmt.Fprintf("%s: %s", program, err.Error())
	os.Exit(1)
}
