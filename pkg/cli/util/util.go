package util

import (
	"fmt"
	"os"
)

func CheckErr(err error) {
	if err == nil {
		return
	}
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}
