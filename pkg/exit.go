package pkg

import (
	"fmt"
	"os"
)

// ExitIfError logs an error and exits 1 if error is not nil
func ExitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
