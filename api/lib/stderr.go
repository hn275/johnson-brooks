package lib

import (
	"fmt"
	"os"
)

func StdErr(e interface{}) {
	stderr := os.Stderr
	if _, err := stderr.WriteString(fmt.Sprintf("[ERROR]: %v\n", e)); err != nil {
		panic(err)
	}
}
