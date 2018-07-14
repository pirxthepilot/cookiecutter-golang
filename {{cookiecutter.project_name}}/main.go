package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "undefined"

// Main
func main() {

	// Flags processing
	versionFlag := flag.Bool("version", false, "Show version")
	flag.Parse()

	// Version
	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

}
