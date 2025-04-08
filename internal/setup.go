package interal

import (
	"fmt"
	"os"
)

func Setup(conflocation string, webroot string) {

	file, ferr := os.Create(conflocation) // Creates the configuration file.
	if ferr != nil {                      // Checks if the file already exists
		fmt.Println("File most likely exists or permission error!", ferr) // Error Checking
	} else {
		defer file.Close() // defer ensures the file closes after no matter what.
	}

	derr := os.Mkdir(webroot, 0755) // Creates a new directroy where the owner can rwx. Non-owner can r-x.
	if derr != nil {
		fmt.Println("Directory could not be made, most likely directory exists!", derr) // Error handling
	}
}
