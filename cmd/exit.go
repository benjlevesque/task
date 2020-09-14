package cmd

import (
	"fmt"
	"os"
)

// exitIfError logs an error and exits 1 if error is not nil
func exitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
