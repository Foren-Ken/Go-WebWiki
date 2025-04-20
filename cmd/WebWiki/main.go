package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	internl "github.com/Foren-Ken/Go-WebWiki/internal"
	server "github.com/Foren-Ken/Go-WebWiki/pkg/server"
)

func main() {

	webrootflag := flag.String("webroot", "webroot", "Defines the root of the webserver")

	flag.Parse()

	dir, wderr := os.Getwd()

	if wderr != nil {
		fmt.Println("Current Directory cannot be found!", wderr)
		os.Exit(1)
	}

	webroot := filepath.Join(dir, *webrootflag) // Location of the webroot (nothing should access the parent)

	conflocation := filepath.Join(dir, "Go-WebWiki.conf") // Location of the Configuratiuon file.

	internl.Setup(conflocation, webroot)

	server.Webroot = webroot
	server.StartServer()
	// fmt.Println(server.FileConfirm("mainn.go", "test"))
}
