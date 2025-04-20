package internal

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
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

	var githuburl string

	fmt.Print("Please input a github wiki directory.\n Example: https://github.com/Foren-Ken/tech-journal.wiki.git\n")
	fmt.Scanln(&githuburl)

	_, err := git.PlainClone(webroot, false, &git.CloneOptions{
		URL:      githuburl,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Printf("Error Cloning Repository %s\n", err)
		return
	}
	fmt.Printf("Sucessfully Cloned Repository %s\n", githuburl)
}
